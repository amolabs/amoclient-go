package tx

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/util"
)

var GrantCmd = &cobra.Command{
	Use:   "grant <parcel_id> <address> <key_custody>",
	Short: "Grant a parcel permission",
	Args:  cobra.MinimumNArgs(3),
	RunE:  grantFunc,
}

func grantFunc(cmd *cobra.Command, args []string) error {
	parcel, err := hex.DecodeString(args[0])
	if err != nil {
		return err
	}

	grantee, err := hex.DecodeString(args[1])
	if err != nil {
		return err
	}

	custody, err := hex.DecodeString(args[2])
	if err != nil {
		return err
	}

	key, err := GetRawKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	result, err := rpc.Grant(parcel, grantee, custody, key)
	if err != nil {
		return err
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		return err
	}

	fmt.Println(string(resultJSON))

	return nil
}
