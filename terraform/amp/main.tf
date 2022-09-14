
module "config" {
  source = "git::https://github.eagleview.com/engineering/amp-config.git//${platform_choice}${test_config_branch}"
}

module "batch" {
  source = "git::https://github.eagleview.com/infrastructure/terraform-cloudops-module-aws-batch.git//batch/?ref=CL-93"

  for_each = module.config.batch_configmap

  resource_name_prefix     = local.resource_name_prefix
  batch_name               = each.key
  container_properties     = each.value.container_properties
  compute_environments     = each.value.compute_environments
}

data "aws_caller_identity" "current" {}

// Useful to troubleshoot role issues
output "assumed-identity-arn" {
  value = data.aws_caller_identity.current.arn
}
