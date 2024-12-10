variable "AWS_REGION" {}
variable "AWS_ACCOUNT_ID" {}


provider "aws" {
  region  = var.AWS_REGION
  profile = "default"
}

# EKS Cluster
resource "aws_eks_cluster" "simple_bank" {
  name     = "simple-bank"
  role_arn = "arn:aws:iam::${var.AWS_ACCOUNT_ID}:role/AWSEKSClusterRole"  # Correct string interpolation
  version  = "1.31"

  vpc_config {
    subnet_ids = [
      "subnet-05a719fc592d9488f",
      "subnet-0cbd7697da16bb104",
      "subnet-09804b2c719227698"
    ]
    endpoint_public_access  = true
    endpoint_private_access = true
  }
}

# VPC CNI Add-on
resource "aws_eks_addon" "vpc_cni" {
  cluster_name  = aws_eks_cluster.simple_bank.name
  addon_name    = "vpc-cni"
  addon_version = "v1.19.0-eksbuild.1"
}

# CoreDNS Add-on
resource "aws_eks_addon" "coredns" {
  cluster_name  = aws_eks_cluster.simple_bank.name
  addon_name    = "coredns"
  addon_version = "v1.11.3-eksbuild.2"
}

# Kube Proxy Add-on
resource "aws_eks_addon" "kube_proxy" {
  cluster_name  = aws_eks_cluster.simple_bank.name
  addon_name    = "kube-proxy"
  addon_version = "v1.31.2-eksbuild.3"
}

# EKS Node Group
resource "aws_eks_node_group" "simple_bank" {
  cluster_name    = aws_eks_cluster.simple_bank.name
  node_group_name = "simple-bank"
  node_role_arn   = "arn:aws:iam::${var.AWS_ACCOUNT_ID}:role/AWSEKSNodeRole"  # Correct string interpolation

  subnet_ids = [
    "subnet-05a719fc592d9488f",
    "subnet-0cbd7697da16bb104",
    "subnet-09804b2c719227698"
  ]

  scaling_config {
    desired_size = 1
    min_size     = 0
    max_size     = 2
  }

  instance_types = ["t3.small"]
  disk_size      = 10
  ami_type       = "AL2_x86_64"
}
