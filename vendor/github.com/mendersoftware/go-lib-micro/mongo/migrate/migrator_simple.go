// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package migrate

import (
	"context"
	"sort"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"

	"github.com/mendersoftware/go-lib-micro/log"
)

// SimpleMigratior applies migrations by comparing `Version` of migrations
// passed to Apply() and already applied migrations. Only migrations that are of
// version higher than the last applied migration will be run. For example:
//
//   already applied migrations: 1.0.0, 1.0.1, 1.0.2
//   migrations in Apply(): 1.0.1, 1.0.3, 1.1.0
//   migrations that will be applied: 1.0.3, 1.1.0
//
type SimpleMigrator struct {
	Session *mgo.Session
	Db      string
}

// Apply will apply migrations. After each successful migration a new migration
// record will be added to DB with the version of migration that was just
// applied. If a migration fails, Apply() returns an error and does not add a
// migration record (so last migration that is recorded is N-1).
//
// Apply() will log some messages when running. Logger will be extracted from
// context using go-lib-micro/log.LoggerContextKey as key.
func (m *SimpleMigrator) Apply(ctx context.Context, target Version, migrations []Migration) error {
	l := log.FromContext(ctx).F(log.Ctx{"db": m.Db})

	sort.Slice(migrations, func(i int, j int) bool {
		return VersionIsLess(migrations[i].Version(), migrations[j].Version())
	})

	applied, err := GetMigrationInfo(m.Session, m.Db)
	if err != nil {
		return errors.Wrap(err, "failed to list applied migrations")
	}

	// starts at 0.0.0
	last := Version{}

	if len(applied) != 0 {
		// sort applied migrations wrt. version
		sort.Slice(applied, func(i int, j int) bool {
			return VersionIsLess(applied[i].Version, applied[j].Version)
		})
		// last version from already applied migrations
		last = applied[len(applied)-1].Version
	}

	// try to apply migrations
	for _, migration := range migrations {
		mv := migration.Version()
		if VersionIsLess(target, mv) {
			l.Warnf("migration to version %s skipped, target version %s is lower",
				mv, target)
		} else if VersionIsLess(last, mv) {
			// log, migration applied
			l.Infof("applying migration from version %s to %s",
				last, mv)

			// apply migration
			if err := migration.Up(last); err != nil {
				l.Errorf("migration from %s to %s failed: %s",
					last, mv, err)

				// migration from last to migration.Version() failed: err
				return errors.Wrapf(err,
					"failed to apply migration from %s to %s",
					last, mv)
			}

			if err := UpdateMigrationInfo(mv, m.Session, m.Db); err != nil {

				return errors.Wrapf(err,
					"failed to record migration from %s to %s",
					last, mv)

			}
			last = mv
		} else {
			// log migration already applied
			l.Infof("migration to version %s skipped", mv)
		}
	}

	// ideally, when all migrations have completed, DB should be in `target` version
	if VersionIsLess(last, target) {
		l.Warnf("last migration to version %s did not produce target version %s",
			last, target)
		// record DB version anyways
		if err := UpdateMigrationInfo(target, m.Session, m.Db); err != nil {
			return errors.Wrapf(err,
				"failed to record migration from %s to %s",
				last, target)
		}
	} else {
		l.Infof("DB migrated to version %s", target)
	}

	return nil
}
