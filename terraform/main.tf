resource "aws_lambda_function" "dupyeeter_lambda" {
  function_name = "dupyeeter"
  role          = aws_iam_role.dupyeeter-runtime.arn
  package_type = "Image"
  image_uri = "613475857488.dkr.ecr.us-east-1.amazonaws.com/dupyeeter:latest"
  image_config {
    entry_point = ["/main"]
  }

  environment {
    variables = {
      foo = "bar"
    }
  }
}

resource "aws_ecr_repository" "repository" {
  name = "dupyeeter"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
}
