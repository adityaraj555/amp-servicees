resource "aws_iam_role" "main" {
    count = length(module.config.role_config_map)
    assume_role_policy = <<POLICY
${module.config.role_config_map[count.index].assume_role_policy}
    POLICY
    
    inline_policy {
        name   = "${module.config.role_config_map[count.index].inline_policy_name}"
        policy = <<POLICY
${module.config.role_config_map[count.index].inline_policy}
        POLICY
    }
    max_session_duration = "3600"
    name                 = "${local.resource_name_prefix}-role-${module.config.role_config_map[count.index].amp_role_name}"
    path                 = "/"
}