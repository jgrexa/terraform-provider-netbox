package util

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func RenderTemplate(tplstring string, data map[string]string) string {
	tmpl, err := template.New("test").Parse(tplstring)
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer

	// tmpl.Execute(os.Stdout, data)
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		panic(err)
	}
	return tpl.String()
}

func TestAccResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID set")
		}

		return nil
	}
}

func TestAccPreCheck(t *testing.T) {
	if err := os.Getenv("NETBOX_TOKEN"); err == "" {
		t.Fatal("NETBOX_TOKEN must be set for acceptance tests")
	}
}
