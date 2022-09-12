
module "config" {
  source = "git::https://github.eagleview.com/engineering/amp-config.git//${platform_choice}${test_config_branch}"
}

module "" {
  source = "git::https://github.eagleview.com/infrastructure/terraform-cloudops-module-aws-batch.git//batch/?ref=0.0.1"

  for_each = module.config.batch_configmap

  resource_name_prefix  = local.resource_name_prefix
  vpc_id                = each.value.vpc_id
  subnet_id             = each.value.subnet_id
  batch_name            = each.key
  container_properties  = each.value.container_properties
  compute_environment_name = each.value.compute_environment_name
  instance_type         = each.value.instance_type
}
data "aws_caller_identity" "current" {}

// Useful to troubleshoot role issues
output "assumed-identity-arn" {
  value = data.aws_caller_identity.current.arn
}
