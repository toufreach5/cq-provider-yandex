// Code generated by yandex cloud generator; DO NOT EDIT.

package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func ComputeImages() *schema.Table {
	return &schema.Table{
		Name:         "yandex_compute_images",
		Resolver:     fetchComputeImages,
		Multiplex:    client.FolderMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteFolderFilter,
		Columns: []schema.Column{
			{
				Name:        "id",
				Type:        schema.TypeString,
				Description: "ID of the image.",
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "folder_id",
				Type:        schema.TypeString,
				Description: "ID of the folder that the image belongs to.",
				Resolver:    client.ResolveFolderID,
			},
			{
				Name:        "created_at",
				Type:        schema.TypeTimestamp,
				Description: "",
				Resolver:    client.ResolveAsTime,
			},
			{
				Name:        "name",
				Type:        schema.TypeString,
				Description: "Name of the image. 1-63 characters long.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the image. 0-256 characters long.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "labels",
				Type:        schema.TypeJSON,
				Description: "Resource labels as `key:value` pairs. Maximum of 64 per resource.",
				Resolver:    client.ResolveLabels,
			},
			{
				Name:        "family",
				Type:        schema.TypeString,
				Description: "The name of the image family to which this image belongs.\n\n You can get the most recent image from a family by using\n the [yandex.cloud.compute.v1.ImageService.GetLatestByFamily] request\n and create the disk from this image.",
				Resolver:    schema.PathResolver("Family"),
			},
			{
				Name:        "storage_size",
				Type:        schema.TypeBigInt,
				Description: "The size of the image, specified in bytes.",
				Resolver:    schema.PathResolver("StorageSize"),
			},
			{
				Name:        "min_disk_size",
				Type:        schema.TypeBigInt,
				Description: "Minimum size of the disk which will be created from this image.",
				Resolver:    schema.PathResolver("MinDiskSize"),
			},
			{
				Name:        "product_ids",
				Type:        schema.TypeStringArray,
				Description: "License IDs that indicate which licenses are attached to this resource.\n License IDs are used to calculate additional charges for the use of the virtual machine.\n\n The correct license ID is generated by Yandex.Cloud. IDs are inherited by new resources created from this resource.\n\n If you know the license IDs, specify them when you create the image.\n For example, if you create a disk image using a third-party utility and load it into Yandex Object Storage, the license IDs will be lost.\n You can specify them in the [yandex.cloud.compute.v1.ImageService.Create] request.",
				Resolver:    schema.PathResolver("ProductIds"),
			},
			{
				Name:        "status",
				Type:        schema.TypeString,
				Description: "Current status of the image.",
				Resolver:    client.EnumPathResolver("Status"),
			},
			{
				Name:        "os_type",
				Type:        schema.TypeString,
				Description: "Operating system type. The default is `LINUX`.\n\n This field is used to correctly emulate a vCPU and calculate the cost of using an instance.",
				Resolver:    client.EnumPathResolver("Os.Type"),
			},
		},
	}

}

func fetchComputeImages(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &compute.ListImagesRequest{FolderId: f}
		it := c.Services.Compute.Image().ImageIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}

	return nil
}
