resource "aws_iam_role" "sim2pdw" {
  assume_role_policy = <<POLICY
${module.config.environment_config_map.trust_relashionships_external_service}
  POLICY



  managed_policy_arns = []

  max_session_duration = "3600"
  name                 = "${local.resource_name_prefix}-role-sim2pdw"
  path                 = "/"

  tags = {
    Name        = "${local.resource_name_prefix}-role-sim2pdw"
    Description = "AWS IAM role to allow services to access platform-data-orchestrator common resources like Lambda"
  }
}

resource "aws_iam_role" "querypdw" {
  assume_role_policy = <<POLICY
${module.config.environment_config_map.trust_relashionships_external_service}
  POLICY



  managed_policy_arns = []

  max_session_duration = "3600"
  name                 = "${local.resource_name_prefix}-role-querypdw"
  path                 = "/"

  tags = {
    Name        = "${local.resource_name_prefix}-role-querypdw"
    Description = "AWS IAM role to allow services to access platform-data-orchestrator common resources like Lambda"
  }
}

resource "aws_iam_role" "kafkapublisher" {
  assume_role_policy = <<POLICY
${module.config.environment_config_map.trust_relashionships_external_service}
  POLICY


  managed_policy_arns = []

  max_session_duration = "3600"
  name                 = "${local.resource_name_prefix}-role-kafkapublisher"
  path                 = "/"

  tags = {
    Name        = "${local.resource_name_prefix}-role-kafkapublisher"
    Description = "AWS IAM role to allow services to access platform-data-orchestrator common resources like Lambda"
  }
}
