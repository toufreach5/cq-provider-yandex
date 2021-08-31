// Code generated by yandex cloud generator; DO NOT EDIT.

package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

func ResourceManagerClouds() *schema.Table {
	return &schema.Table{
		Name:        "yandex_resourcemanager_clouds",
		Resolver:    fetchResourceManagerClouds,
		Multiplex:   client.MultiplexBy(client.Clouds),
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:            "id",
				Type:            schema.TypeString,
				Description:     "ID of the cloud.",
				Resolver:        client.ResolveResourceId,
				CreationOptions: schema.ColumnCreationOptions{Nullable: false, Unique: true},
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
				Description: "Name of the cloud. 3-63 characters long.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the cloud. 0-256 characters long.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "organization_id",
				Type:        schema.TypeString,
				Description: "ID of the organization that the cloud belongs to.",
				Resolver:    schema.PathResolver("OrganizationId"),
			},
		},
	}

}

func fetchResourceManagerClouds(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	cloud, err := c.Services.ResourceManager.Cloud().Get(ctx, &resourcemanager.GetCloudRequest{CloudId: c.MultiplexedResourceId})
	if err != nil {
		return err
	}

	res <- cloud

	return nil
}
