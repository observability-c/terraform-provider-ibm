---

subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM : Cloud Object Storage Lifecycle Configuration"
description: 
  "Manages IBM Cloud Object Storage Lifecycle Configuration"
---

# ibm_cos_bucket_lifecycle_configuration
Provides the recommended way of managing the lifecycle configuration for a bucket. This provides an independent resource to manage the lifecycle configuration for a bucket. A lifecycle configuration includes one or more lifecycle rules. Each rule has an unique id, filter, status and an action. There are 4 kinds of actions:  [transition](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-archive), [expiration](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-expiry), noncurrent version expiration and abort incomplete multipart upload. A lifecycle configuration can have multiple expiration rules but only one transition rule. A given bucket can only have one ibm_cos_bucket_lifecycle_configuration resource.


# Example usage

## Adding lifecycle configuration with expiration.

Adding lifecycle configuration with expiration and prefix filter.

```terraform


resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class

}
resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
  bucket_crn = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  lifecycle_rule {
    expiration{
      days = 1
    }
    filter {
      prefix = "foo"
    }  
    rule_id = "id"
    status = "enable"
  
  }
}

```

## Adding lifecycle configuration with transition.

Adding lifecycle configuration with transition.

```terraform


resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class

}
resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
  bucket_crn = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  lifecycle_rule {
    transition{
      days = 1
      storage_class = "GLACIER"
    }
    filter {
      prefix = ""
    }  
    rule_id = "id"
    status = "enable"
  
  }
}

```
## Adding lifecycle configuration with abort incomplete multipart upload.

```terraform


resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class

}
resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
  bucket_crn = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  lifecycle_rule {
    abort_incomplete_multipart_upload{
      days_after_initiation = 1
    }
    filter {
      prefix = ""
    }  
    rule_id = "id"
    status = "enable"
  
  }
}

```
## Adding lifecycle configuration with non-current version expiration.

```terraform


resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class

}
resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
  bucket_crn = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  lifecycle_rule {
    noncurrent_version_expiration{
			   noncurrent_days = "1
		}
    filter {
      prefix = ""
    }  
    rule_id = "id"
    status = "enable"
  
  }
}

```
## Adding lifecycle configuration with multiple rules.

```terraform


resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class

}
resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
  bucket_crn = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  lifecycle_rule {
    expiration{
      days = 1
    }
    filter {
      prefix = "foo"
    }  
    rule_id = "id"
    status = "enable"
  }
    lifecycle_rule {
    expiration{
      days = 2
    }
    filter {
      prefix = "bar"
    }  
    rule_id = "id2"
    status = "enable"
  }
}

```

## Adding lifecycle configuration for object expiration with filter based on object size.

```terraform


resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class

}
resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
  bucket_crn = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  lifecycle_rule {
    expiration{
      days = 1
    }
    filter {
      object_size_greater_than = 1000
    }  
    rule_id = "id"
    status = "enable"
  }
}

```

## Adding lifecycle configuration for object expiration with filter based on tag.

```terraform


resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class

}
resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
  bucket_crn = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  lifecycle_rule {
    expiration{
      days = 1
    }
    filter {
      tag{
        key = "MyObjectTagKey"
        value = "MyObjectTagValue"
      }
    }  
    rule_id = "id"
    status = "enable"
  }
}

```
## Adding lifecycle configuration for object expiration with multiple filter.

```terraform


resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class

}
resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
  bucket_crn = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  lifecycle_rule {
    expiration{
      days = 1
    }
    filter {
      and{
        prefix = "foo"
	tags{
          key = "MyObjectTagKey"
          value = "MyObjectTagValue"
	}
	object_size_greater_than = 20
	object_size_less_than = 40
      }
    }  
    rule_id = "id"
    status = "enable"
  }
}

```

## Adding lifecycle configuration for object expiration with filter based on multiple tags.

```terraform


resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class

}
resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
  bucket_crn = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  lifecycle_rule {
    expiration{
      days = 1
    }
    filter {
      and{
	tags{
          key = "MyObjectTagKey1"
          value = "MyObjectTagValue1"
	}
	tags{
          key = "MyObjectTagKey2"
          value = "MyObjectTagValue2"
	}
      }
    }  
    rule_id = "id"
    status = "enable"
  }
}

```

# Switching from legacy lifecycle rules to ibm_cos_bucket_lifecycle_configuration :


**Note:**
If you use legacy `expire_rule`, `archive_rule`, `noncurrent_version_expiration`, `abort_incomplete_multipart_upload_days` lifecycle rule features on an ibm_cos_bucket, Terraform will assume management over the full set of Lifecycle rules for the bucket, treating additional Lifecycle rules as drift. For this reason, legacy rules cannot be mixed with the external ibm_cos_bucket_lifecycle_configuration resource for a given cos bucket. Users that want to continue using the legacy `expire_rule`, `archive_rule`, `noncurrent_version_expiration`, `abort_incomplete_multipart_upload_days` lifecycle rule features cannot use all of the filter capabilities available with an ibm_cos_bucket_lifecycle_configuration resource.  Also,  using the legacy feature one cannot create a single rule with all the actions of lifecycle configuration. 

In case you want to switch from using legacy lifecycle rules in the definition of an existing bucket to using a bucket lifecycle configuration resource for the existing bucket, please follow the steps below

## Example usage
For example if there is a cos bucket existing in the `.tfstate` file with `expire_rule` applied.
```terraform

# Existing Cos bucket configuration 

resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class
  expire_rule {
    rule_id = "id"
    enable  = true
    days    = 45
    prefix  = "foo"
  }

}

```
Step 1 : Add new `ibm_cos_bucket_lifecycle_configuration` with same configuration as existing.

```terraform

resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class
  expire_rule {
    rule_id = "id"
    enable  = true
    days    = 45
    prefix  = "foo"
  }
}

resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
  bucket_crn = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  lifecycle_rule {
    expiration{
      days = 45
    }
    filter {
      prefix = "foo"
    }  
    rule_id = "id"
    status = "enable"
  
  }
  }
}

```
Step 2 : Run `terraform apply`

Step 3 : Remove the old configuration from the `ibm_cos_bucket` block

```terraform

resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class
}

resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
  bucket_crn = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  lifecycle_rule {
    expiration{
      days = 45
    }
    filter {
      prefix = "foo"
    }  
    rule_id = "id"
    status = "enable"
  
  }
  }
}

```

## Argument reference
Review the argument references that you can specify for your resource. 
- `bucket_crn` - (Required, Forces new resource, String) The CRN of the COS bucket.
- `bucket_location` - (Required, Forces new resource, String) The location of the COS bucket.
- `endpoint_type`- (Optional, String) The type of the endpoint either `public` or `private` or `direct` to be used for buckets. Default value is `public`.
- `lifecycle_rule`- (Required, List) Nested block have the following structure:
  
  Nested scheme for `lifecycle_rule`:
  - `expiration`- (Optional) Configuration block that specifies the expiration for the lifecycle of the object in the form of date, days and, whether the object has a delete marker.
  - `transition`- (Optional) Configuration block that specifies the transition for the object.
  - `abort_incomplete_multipart_upload`- (Optional) Configuration block that specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload.
  - `noncurrent_version`- (Optional) Configuration block that specifies when noncurrent object versions expire.
  - `id`- (Required) Unique id for the rule.
  - `filter`- (Required)  Configuration block used to identify objects that a Lifecycle Rule applies to. If not specified, the rule will be applied to all the objects in a bucket.
  - `status`- (Required) Whether the rule is currently being applied. Valid values: enable or disable.

 Nested scheme for `expiration`:
  - `days`- (Optional)  Days, of the objects that are subject to the rule. The value must be a non-zero positive integer.
  - `date`- (Optional) Date the object is to be deleted. The date value must be in RFC3339 full-date format.
  - `expired_object_delete_marker`- (Optional, Conflicts with date and days) Indicates whether ibm will remove a delete marker with no noncurrent versions.
  
 Nested scheme for `transition`:
  - `days`- (Optional)  Number of days after creation when objects are transitioned to the specified storage class. The value must be a positive integer. If both days and date are not specified, defaults to 0.
  - `date`- (Optional) Date objects are transitioned to the specified storage class. The date value must be in RFC3339 full-date format.
  - `storage_class`- (Optional) Class of storage used to store the object. Valid Values: GLACIER,ACCELERATED.
  
 Nested scheme for `noncurrent_version_expiration`:
  - `noncurrent_days` - (Optional) Number of days an object is noncurrent before lifecycle action is performed. Must be a positive integer.

 Nested scheme for `abort_incomplete_multipart_upload`:
  - `days_after_initiation` - Number of days after which incomplete multipart uploads are aborted.

 Nested scheme for `filter`:
  - `prefix`- (Optional) Prefix identifying one or more objects to which the rule applies.
  - `tag`- (Optional) Key-value map of resource tags.
  - `object_size_greater_than`- (Optional) minimum object size to which the rule applies. Value must be at least 0 if specified.
  - `object_size_less_than`- (Optional) Maximum object size to which the rule applies. Value must be at least 1 if specified.
 

 ## Import IBM COS lifecycle configuration
The lifecycle configuration rules for a bucket can be imported into an `ibm_cos_bucket_lifecycle_configuration` resource for a particular bucket using the bucket id.

id = `$CRN:meta:$buckettype:$bucketlocation`

**Syntax**

```
$ terraform import ibm_cos_bucket_lifecycle_configuration.lifecycle `$CRN:meta:$buckettype:$bucketlocation`

```

**Example**

```

$ terraform import ibm_cos_bucket_lifecycle_configuration.lifecycle crn:v1:bluemix:public:cloud-object-storage:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3:bucket:mybucketname:meta:crl:eu:public

```