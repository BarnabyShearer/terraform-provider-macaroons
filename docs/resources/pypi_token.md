# Password Resource

Provides a method of restricting user-scoped pypi.org API tokens into projects-scoped tokens.

## Example Usage

```hcl
resource "macaroons_pypi_token" "project_token" {
  source_token = "pypi-ABC…"
  project      = "foobar"
}
```

## Argument Reference

* `password` - (Required) A user-scoped API token from pypi.org.
* `project` - (Required) The name of project to create a project-scoped token for.

## Attribute Reference

* `token` - The resulting API token that can only be used for given project.
