data "aws_iam_policy_document" "role-assume-role-policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "dupyeeter-runtime" {
  name               = "dupyeeter-runtime"
  assume_role_policy = data.aws_iam_policy_document.role-assume-role-policy.json
}
