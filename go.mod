// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/littlegirlpppp/fabric-sdk-go-gm

replace github.com/cloudflare/cfssl => ./third_party/github.com/cloudflare/cfssl
require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/cloudflare/cfssl v1.4.1
	github.com/go-kit/kit v0.8.0
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.3.3
	github.com/hyperledger/fabric-config v0.0.5
	github.com/hyperledger/fabric-lib-go v1.0.0
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/miekg/pkcs11 v1.0.3
	github.com/mitchellh/mapstructure v1.3.2
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.7.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.6.1
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
	google.golang.org/grpc v1.29.1
	gopkg.in/yaml.v2 v2.3.0
)

go 1.14
