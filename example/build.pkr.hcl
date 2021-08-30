packer {
  required_plugins {
    salt = {
      version = ">=v1.0.0"
      source  = "github.com/hashicorp/salt"
    }
  }
}

source "happycloud "bar-example" {
  mock = local.bar
}

build {

  source "source.happycloud.bar-example" {
    name = "bar"
  }

  provisioner "salt-masterless" {
    local_state_tree = "/Users/me/salt"
  }
}
