package cli

import (
	"testing"

	"github.com/libopenstorage/openstorage/api"
	"github.com/stretchr/testify/require"
)

func TestCmdMarshalProto(t *testing.T) {
	// TODO @ggriffiths fix after jsonpb work
	t.Skip()
	volumeSpec := &api.VolumeSpec{
		Size:   64,
		Format: api.FSType_FS_TYPE_EXT4,
	}
	data := cmdMarshalProto(volumeSpec, false)
	require.Equal(
		t,
		`{
 "ephemeral": false,
 "size": "64",
 "format": "ext4",
 "block_size": "0",
 "ha_level": "0",
 "cos": "none",
 "io_profile": "sequential",
 "dedupe": false,
 "snapshot_interval": 0,
 "shared": false,
 "aggregation_level": 0,
 "encrypted": false,
 "passphrase": "",
 "snapshot_schedule": "",
 "scale": 0,
 "sticky": false,
 "group_enforced": false,
 "compressed": false,
 "cascaded": false,
 "journal": false,
 "sharedv4": false,
 "queue_depth": 0,
 "force_unsupported_fs_type": false,
 "nodiscard": false,
 "storage_policy": "",
 "fp_preference": false,
 "xattr": "UNSPECIFIED",
 "proxy_write": false,
 "auto_fstrim": false,
 "number_of_chunks": 0
}`,
		data,
	)
}
