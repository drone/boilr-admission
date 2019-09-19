// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"errors"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/admission"
)

// ErrAccessDenied is returned if access is denied.
var ErrAccessDenied = errors.New("admission: access denied")

// New returns a new admission plugin.
func New(param1, param2 string) admission.Plugin {
	return &plugin{
		// TODO replace or remove these configuration
		// parameters. They are for demo purposes only.
		param1: param1,
		param2: param2,
	}
}

type plugin struct {
	// TODO replace or remove these configuration
	// parameters. They are for demo purposes only.
	param1 string
	param2 string
}

func (p *plugin) Admit(ctx context.Context, req *admission.Request) (*drone.User, error) {
	// this demonstrates how we can grant access to allowed
	// users by returning nil.
	switch req.User.Login {
	case "wolverine", "storm", "rogue", "gambit", "cyclops":
		return nil, nil
	}

	// this demonstrates that we can grant administrative
	// access by setting the administrator flag to true
	// and returning the updated user.
	switch req.User.Login {
	case "professorx":
		req.User.Admin = true
		return &req.User, nil
	}

	// else we deny access to all other users.
	return nil, ErrAccessDenied
}
