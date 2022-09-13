
module "config" {
  source = "git::https://github.eagleview.com/engineering/amp-config.git//${platform_choice}${test_config_branch}"
}

data "aws_caller_identity" "current" {}

// Useful to troubleshoot role issues
output "assumed-identity-arn" {
  value = data.aws_caller_identity.current.arn
}
