terraform {
  backend "s3" {
    bucket = "dupyeeter-infra"
    key    = "tf.tfstate"
    region = "us-east-1"
  }
}

