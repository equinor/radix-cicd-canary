package privateimagehub

import (
	"context"
	"math/rand"
	"time"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/privateimagehub"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Set runs tests related to private image hub. Expect canary2 to be built and deployed before test run
func Set(ctx context.Context, cfg config.Config) error {
	// Due to a timing bug in Config Syncer (https://github.com/kubeops/config-syncer) that can happen
	// when a new namespace is created and at the same time a secret that must be synced to the namespace is updated,
	// the old "cached" secret from the nsSyncer overwrites the secret created by the secret informer's OnUpdate.
	// Currently, this timing bug is activly happening in the playground cluster almost every time this suite is executed.
	// The random sleep (5-10 sec) will allow the nsSyncer in Config Syncer to update perform the initial sync before we update the secret.
	time.Sleep(time.Duration(rand.Intn(10)+5) * time.Second)

	log.Ctx(ctx).Debug().Str("app", defaults.App3Name).Msg("set privateimagehub passford for application")
	err := privateimagehub.SetPassword(cfg, defaults.App3Name)
	if err != nil {
		return errors.Errorf("failed to set private image hub password. %v", err)
	}

	return nil
}
