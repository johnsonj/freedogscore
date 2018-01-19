provider "google" {
  version = "v1.5.0"

  region  = "us-central1"
  project = "${var.project}"
}

provider "random" {
  version = "v1.1.0"
}

provider "archive" {}

variable "project" {
  description = "GCP Project hosting FreeDogScore"
}

// HACK: Used to get the md5 of function.zip which is used as the object name to force
// the function to redeploy on change
data "archive_file" "function" {
  source_file = "../function.zip"
  output_path = "function.zip"
  type        = "zip"
}

resource "random_string" "bucket_suffix" {
  length  = 8
  special = false
  upper   = false
}

resource "google_storage_bucket" "bucket" {
  name = "freedogscore-${random_string.bucket_suffix.result}"
}

resource "google_storage_bucket_object" "archive" {
  name   = "function-${data.archive_file.function.output_md5}.zip"
  bucket = "${google_storage_bucket.bucket.name}"
  source = "../function.zip"
}

resource "google_cloudfunctions_function" "upload" {
  name                  = "upload"
  description           = "freedogscore upload"
  available_memory_mb   = 128
  source_archive_bucket = "${google_storage_bucket.bucket.name}"
  source_archive_object = "${google_storage_bucket_object.archive.name}"
  trigger_http          = true
  timeout               = 60
  entry_point           = "helloWorld"
}

data "google_cloudfunctions_function" "function_http" {
  name = "${google_cloudfunctions_function.upload.name}"
}

output "upload_url" {
  value = "${google_cloudfunctions_function.upload.https_trigger_url}"
}
