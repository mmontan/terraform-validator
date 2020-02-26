// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

func GetIdentityPlatformTenantDefaultSupportedIdpConfigCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//identityplatform.googleapis.com/projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetIdentityPlatformTenantDefaultSupportedIdpConfigApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "identityplatform.googleapis.com/TenantDefaultSupportedIdpConfig",
			Resource: &AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/identityplatform/v2/rest",
				DiscoveryName:        "TenantDefaultSupportedIdpConfig",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetIdentityPlatformTenantDefaultSupportedIdpConfigApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	clientIdProp, err := expandIdentityPlatformTenantDefaultSupportedIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("client_id"); !isEmptyValue(reflect.ValueOf(clientIdProp)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformTenantDefaultSupportedIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("client_secret"); !isEmptyValue(reflect.ValueOf(clientSecretProp)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}
	enabledProp, err := expandIdentityPlatformTenantDefaultSupportedIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	return obj, nil
}

func expandIdentityPlatformTenantDefaultSupportedIdpConfigClientId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantDefaultSupportedIdpConfigClientSecret(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantDefaultSupportedIdpConfigEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
