package aws

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAWSOrganizationAccountImportBasic(t *testing.T) {
	resourceName := "aws_organization_account.test"
	name := "my_new_account"
	email := "foo@bar.org"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSOrganizationAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSOrganizationAccountConfig(name, email),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
