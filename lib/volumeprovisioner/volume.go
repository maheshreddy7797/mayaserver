// This file abstracts & exposes persistent volume provisioner features. All
// maya api server's persistent volume provisioners need to implement these
// contracts.
package volumeprovisioner

import (
	"github.com/openebs/mayaserver/lib/api/v1"
)

// VolumeInterface abstracts the persistent volume features of any persistent
// volume provisioner.
//
// NOTE:
//    maya api server can make use of any persistent volume provisioner & execute
// corresponding volume related operations.
type VolumeInterface interface {
	// Label assigned against this persistent volume provisioner
	Label() string

	// Name of the persistent volume provisioner
	Name() string

	// Profile will set the persistent volume provisioner's profile
	//
	// NOTE:
	//    Will return false if profile is not supported by the persistent
	// volume provisioner.
	//
	// NOTE:
	//    This is used to set the persistent volume provisioner profile lazily
	// i.e. much after the initialization of persistent volume provisioner instance.
	// It is assumed that persistent volume claim will be available at the time of
	// invoking this method.
	Profile(*v1.PersistentVolumeClaim) (bool, error)

	// TODO
	// Deprecate in favour of Adder
	//
	// Provisioner gets the instance capable of provisioning volumes w.r.t this
	// persistent volume provisioner.
	//
	// Note:
	//    Will return false if provisioning of volumes is not supported by the
	// persistent volume provisioner.
	Provisioner() (Provisioner, bool)

	// Deleter gets the instance capable of deleting volumes w.r.t this
	// persistent volume provisioner.
	//
	// Note:
	//    Will return false if deletion of volumes is not supported by the
	// persistent volume provisioner.
	Deleter() (Deleter, bool)

	// TODO
	// Deprecate in favour of Reader
	//
	// Informer gets the instance capable of providing volume information w.r.t this
	// persistent volume provisioner.
	//
	// Note:
	//    Will return false if providing volume information is not supported by
	// the persistent volume provisioner.
	Informer() (Informer, bool)

	// Reader gets the instance capable of providing persistent volume information
	// w.r.t this persistent volume provisioner.
	//
	// Note:
	//    Will return false if providing persistent volume information is not
	// supported by this persistent volume provisioner.
	Reader() (Reader, bool)

	// Adder gets the instance capable of creating a persistent volume
	// w.r.t this persistent volume provisioner.
	//
	// Note:
	//    Will return false if creating persistent volume is not
	// supported by this persistent volume provisioner.
	Adder() (Adder, bool)
}

// TODO
// Deprecate in favour of Reader
//
// Informer interface abstracts fetching of volume related information
// from a persistent volume provisioner.
type Informer interface {
	// Info tries to fetch the volume details from the persistent volume
	// provisioner.
	Info(*v1.PersistentVolumeClaim) (*v1.PersistentVolume, error)
}

// Reader interface abstracts fetching of persistent volume related information
// from a persistent volume provisioner.
type Reader interface {
	// Read fetches the volume details from the persistent volume
	// provisioner.
	Read(*v1.PersistentVolumeClaim) (*v1.PersistentVolumeList, error)
}

// TODO
// Deprecate in favour of Adder
//
// Provisioner interface abstracts creation of volume from a persistent volume
// provisioner.
type Provisioner interface {
	// Provision tries to create a volume of a persistent volume provisioner.
	Provision(*v1.PersistentVolumeClaim) (*v1.PersistentVolume, error)
}

// Adder interface abstracts creation of persistent volume from a persistent
// volume provisioner.
type Adder interface {
	// Add creates a new persistent volume
	Add(*v1.PersistentVolumeClaim) (*v1.PersistentVolumeList, error)
}

// Deleter interface abstracts deletion of volume of a persistent volume
// provisioner.
type Deleter interface {
	// Delete tries to delete a volume of a persistent volume provisioner.
	Delete(*v1.PersistentVolume) (*v1.PersistentVolume, error)
}
