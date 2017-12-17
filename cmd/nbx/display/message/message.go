package message

import (
  "github.com/renstrom/dedent"  
)

func OverviewDescription() string {
  return dedent.Dedent(`
    The Nanobox platform within your console.

                   **
                ********
             ***************
          *********************
            *****************
          ::    *********    ::
             ::    ***    ::
           ++   :::   :::   ++
              ++   :::   ++
                 ++   ++
                    +
    _  _ ____ _  _ ____ ___  ____ _  _
    |\ | |__| |\ | |  | |__) |  |  \/
    | \| |  | | \| |__| |__) |__| _/\_
  `)
}
