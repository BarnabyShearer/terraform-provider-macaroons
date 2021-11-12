# Macaroons Provider

Macaroons are flexible authorization credentials that support decentralized delegation, attenuation, and verification. Given an existing credential this provider can attenuate it for a specific use.

At the moment it is tested for scoping a pypi.org's API token per project (PRs for other uses welcome).

## Example Usage

```hcl
terraform {
  required_providers {
    macaroons = {
      source = "BarnabyShearer/macaroons"
    }
  }
}

resource "macaroons_pypi_token" "efm8" {
  source_token = "pypi-ABC…"
  project      = "efm8"
}
```
