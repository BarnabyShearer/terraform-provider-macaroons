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
