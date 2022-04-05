package azure

import (
	"time"
)

type statefulNode struct {
	ID                *string      `json:"id,omitempty"`
	Name              *string      `json:"name,omitempty"`
	ResourceGroupName *string      `json:"resourceGroupName,omitempty"`
	Region            *string      `json:"region,omitempty"`
	Strategy          *Strategy    `json:"strategy,omitempty"`
	Compute           *Compute     `json:"compute,omitempty"`
	Persistence       *Persistence `json:"persistence,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// forceSendFields is a list of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string

	// nullFields is a list of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string
}

type Strategy struct {
	DrainingTimeout    *int  `json:"drainingTimeout,omitempty"`
	FallbackToOnDemand *bool `json:"fallbackToOd,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Compute struct {
	OS                  *string              `json:"os,omitempty"`
	VMSizes             *VMSizes             `json:"vmSizes,omitempty"`
	LaunchSpecification *LaunchSpecification `json:"launchSpecification,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type VMSizes struct {
	OnDemandSizes []string `json:"odSizes,omitempty"`
	SpotSizes     []string `json:"spotSizes,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LaunchSpecification struct {
	Image               *Image               `json:"image,omitempty"`
	Network             *Network             `json:"network,omitempty"`
	Login               *Login               `json:"login,omitempty"`
	CustomData          *string              `json:"customData,omitempty"`
	LoadBalancersConfig *LoadBalancersConfig `json:"loadBalancersConfig,omitempty"`
	OSDisk              *OSDisk              `json:"osDisk,omitempty"`
	DataDisks           *DataDisks           `json:"dataDisks,omitempty"`
	BootDiagnostics     *BootDiagnostics     `json:"bootDiagnostics,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Image struct {
	MarketPlace *MarketPlaceImage `json:"marketplace,omitempty"`
	Custom      *CustomImage      `json:"custom,omitempty"`
	Gallery     *Gallery          `json:"gallery,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type MarketPlaceImage struct {
	Publisher *string `json:"publisher,omitempty"`
	Offer     *string `json:"offer,omitempty"`
	SKU       *string `json:"sku,omitempty"`
	Version   *string `json:"version,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type CustomImage struct {
	ResourceGroupName *string `json:"resourceGroupName,omitempty"`
	Name              *string `json:"name,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Gallery struct {
	ResourceGroupName *string `json:"resourceGroupName,omitempty"`
	GalleryName       *string `json:"galleryName,omitempty"`
	ImageName         *string `json:"imageName,omitempty"`
	VersionName       *string `json:"versionName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Network struct {
	VirtualNetworkName *string             `json:"virtualNetworkName,omitempty"`
	ResourceGroupName  *string             `json:"resourceGroupName,omitempty"`
	NetworkInterfaces  []*NetworkInterface `json:"networkInterfaces,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type NetworkInterface struct {
	SubnetName     *string `json:"subnetName,omitempty"`
	AssignPublicIp *bool   `json:"assignPublicIp,omitempty"`
	IsPrimary      *bool   `json:"isPrimary,omitempty"`
	PublicIpSku    *string `json:"publicIpSku,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Login struct {
	UserName     *string `json:"userName,omitempty"`
	SSHPublicKey *string `json:"sshPublicKey,omitempty"`
	Password     *string `json:"password,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LoadBalancersConfig struct {
	LoadBalancers []*LoadBalancer `json:"loadBalancers,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LoadBalancer struct {
	Type              *string  `json:"type,omitempty"`
	ResourceGroupName *string  `json:"resourceGroupName,omitempty"`
	Name              *string  `json:"name,omitempty"`
	SKU               *string  `json:"sku,omitempty"`
	BackendPoolNames  []string `json:"backendPoolNames,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type OSDisk struct {
	SizeGB *int    `json:"sizeGB,omitempty"`
	Type   *string `json:"type,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DataDisks struct {
	SizeGB *int    `json:"sizeGB,omitempty"`
	Lun    *int    `json:"lun,omitempty"`
	Type   *string `json:"type,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type BootDiagnostics struct {
	IsEnabled  *bool   `json:"isEnabled,omitempty"`
	Type       *string `json:"type,omitempty"`
	StorageUrl *string `json:"storageUrl"`
}

type Persistence struct {
	ShouldPersistOsDisk     *bool   `json:"shouldPersistOsDisk,omitempty"`
	OsDiskPersistenceMode   *string `json:"osDiskPersistenceMode"`
	ShouldPersistDataDisk   *bool   `json:"shouldPersistDataDisk,omitempty"`
	DataDiskPersistenceMode *string `json:"dataDiskPersistenceMode"`
	ShouldPersistNetwork    *bool   `json:"shouldPersistNetwork"`
	ShouldPersistVm         *bool   `json:"shouldPersistVm"`
}
