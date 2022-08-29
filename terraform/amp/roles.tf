resource "aws_iam_role" "account-management" {
  assume_role_policy = <<POLICY
${module.config.environment_config_map.trust_relashionships_external_service}
  POLICY

  managed_policy_arns = []

  max_session_duration = "3600"
  name                 = "${local.resource_name_prefix}-role-account-management"
  path                 = "/"

  tags = {
    Name        = "${local.resource_name_prefix}-role-account-management"
    Description = "AWS IAM role to allow services to access platform-data-orchestrator common resources like Lambda"
  }
}

