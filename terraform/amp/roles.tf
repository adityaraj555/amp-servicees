resource "aws_iam_role" "sim2pdw" {

  inline_policy {
    name   = "platform-data-orchestrator-resources-access-policy"
    policy = <<POLICY
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": "lambda.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        }
    ]
}
    POLICY
    
  }

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

  inline_policy {
    name   = "platform-data-orchestrator-resources-access-policy"
    policy = <<POLICY
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": "lambda.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        }
    ]
}
    POLICY
    
  }

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

  inline_policy {
    name   = "platform-data-orchestrator-resources-access-policy"
    policy = <<POLICY
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": "lambda.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        }
    ]
}
    POLICY
    
  }

  managed_policy_arns = []

  max_session_duration = "3600"
  name                 = "${local.resource_name_prefix}-role-kafkapublisher"
  path                 = "/"

  tags = {
    Name        = "${local.resource_name_prefix}-role-kafkapublisher"
    Description = "AWS IAM role to allow services to access platform-data-orchestrator common resources like Lambda"
  }
}
