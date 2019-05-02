// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
	return "eJzsWktz2zgSvudXdOU0MxWrdpLMHnzYqozj2knVPDxjp+bItICWiDUIMHhYlis/fqsBknpRtiyT2svqkHIECv19/UJ3g2dwS8tzwIV/BRBU0HQOr3HhX78CkOSFU3VQ1pzDv14BAHzBhf8ClZVREwirNYng4cPf11BZo4J1ysyhouCU8DBztkprF9pGucAgyskrAEea0NM5zPEVwEyRlv487X4GBitq0fAnLGt+0NlYN9/0gNrcZH0j0Qnulvr23Ltv/nxZbfMFhDUBlfEQSuqYhhIDLMgReOGwJrnF/W/+LSxKJcrVBj0a82QCTJcg1WxGjv/DPHyNgvxkDVOnwikFXPt+Ww/ruuh22lht1XFLy4V1cmvtEaXw56ak1bbgaxJqpkjCoiQDXyO55ZoFAGs16UXWsJ/80IvMTv9DImwt5S+L/IS0caq3aa09UVRY18rMm8df//D6eTR/W7eyoxCdaS18sUEvc+4nKVVFxitrjufZb6MBia6xyTZZR70TXCTeviiqSLw9UTjxDy8v3vZG0PzQ+BF1nAQbUE/qHTNl8l6gJlnMtMXtBw4IpJqcIBNwTmBngFpbgYEkAwdhqzoGgmhUaNSDjkBEx0lCL0EZiJ7AmqRHZXxAI6jfE5mIcCRVKKLHeX8+eBkXE6spOeZxcfUZsjDPCSLbYx0jzKxLT8WgtHpA3vZJ3FPU/NtRkROmMF8nkBVvVthL9IBCuEgSvOJvVIAFetAYjShJgnXgA7pAcj8pH12toy9OSK4RucmsxDuCKZFZWQoNRKNVpdgTO9opv/PPLq4+X6Qdfs6Y4Q51JFAeHsjZQxn7QpTo5rSd2QamnDj1EudYMjZAjUqCtAvD1Hft/wbQyCbthDJ6UEZExzpCKRWjQA2ZSj91Q2Fh3e1EmUmN4paCH5VxIwMcCVJ37IyG80oLA5QJ5GZcW2wH5ePwbQwnxZ/SuI1hKPzKTKbLQOOCTxJGUf2psA+ldqn8rbITRyhHwf5zo2lsygTG2qUqH6wjuLM6VuQB71BpnGqCYA9HvnAq0IjQef9AhjENjj1p3dZDA7+wVa2JD4Wkd1uTSye3P94EXMPgWhNRk1NWsjsGVR1ioDFZJgnrNI+01TEkfcAQ/USUJG6LGSq956DU1syfx+8vqq0Lns/zUJLbRMq1TY3ek4SpDeXmYsYECVM6FXnVL32ganNN5YpUow9QKRPD4SSLvN+JuY5BpJXzP6DSb7FDyXQpRljH/0TT3/nwkTAn9+JGwTryqSF4Or91q6rCOU1Uf0wcPdb49DEFJcPg/bm7khzMubB/Dr5VXzphGww4fvlkpOIeceUJkkLyuPVmWHkgw7loTwfSAa2dusNAE2l8wUvDKrTZHT7+fp0Et+rdqSoORKnqfk/c/voZ0D5d3b3nUt6R94DeW6FSB75QTfp7NtY41UqMpdC0+Y4+D/TKBtqAWmwV1+C45OSiBHy66la+YwV/D1MbjWwPxueqNIXQRFjZr82jE1Had1uHb4D7e/jxn2dTFSAar+Ym9cFJyEFIh7d7L1L4riYjOdy/gYvG5L98GUNQZn6WetpvEMhVyiSf/sYVSxoWtn+S/P4JRqHkYs8XNbmCU/VYR0Ejh4uj7ljYHT/6d4VEpZcFl16bY6znzyK3N9saTKY1aNZGHVNev3vpnH8auWse2PF+bzIMwvW7VsJj0r16GKdjSpmm4kKE4UgMyMVM07ByDS5zif0YxuxphZ0VeXY+BsY0KV7z6UZSD8ZU73Av0bqX0Fy+9Qzc/bvC0ddIPrzU15tt1ry8+eb/vv2I3zQ68vkaYKBivs9Zfrm5ueqkQYUy9ZFo4EOFD9ascL4BR3N0UrcH77Lec4B22OfUX8Ifh3wL878vb7Zws3O3vs9Ov8vhCbx1HBHv1efB8UrStKfIHwTyx8tfL28uh0ZdEg41VOjB/Mvlh48H+fNTvmD9mM7wx/W2NxyF0pPevcgdBucKyfXlr5cXN/BHMjpcWBM40Q7sFZlJ4QUaQ+MMb/sGz+3B3sjN/dfB1F/CtL3cPznV7q2CE3DVaswo6rClCoJlpQpaZOg+H7SP4eReRVuU41shm2AlLwXMYcfuouTSiIk58rU1nlsxoaMkbhinVu55CyTWp6TWSsu2aMo0wK7YY5xvnp/pyDnr/OT9/f14bvT+/h6EVuztSVw3v7SSDrJRjiRsbqTtDEilEdk/wDr48VFiP41J7Kf7e/Dk7sidkJjGQEYsJzPlfCjYOSZVv/cdx7Emd9Y6VVAV5WYhx32+Qln5HHF70N2L5ncMdjgGm18y2Iiw9CJFunqZUpcxHyecKuu23RmWM2msfb7C2cM9aTuF4opvM/ZKk/e0khqkp6Kv6wC/+pe1fl/9CV+5uv7z+qUNn9WSfCgq8h7nVOCcJp7EgFbEunb2XlUYCJpXr1gtWS4Ya85yQS+hwdDenXyNFPf0Ws2TqRfA5WC3dTeb2aSVsgFo9T5LIztdzhkb1q4k8ymHaVqoqoqkwkB6z4HVcTE2FHfKq903LYdJjh2djoEyMNNqXu45hTpkJ0G1rb7gFN2hXgX7gf7ArjQu0tZfn4WszU/jQuuq3OkSBGrt23T4Vxb/WxNiKPa/BdhB5kwzss2lzBkbH9MhVXVYFo0ChzxgVoi21PPh6lOrPo4VqXKEZ+0CtgT23E+TWeXTJ8a0Pbhn1lUYzqHvR4dcXqgHOkLHeWnYsd/1n9dNzkz7/jcAAP//AykX4A=="
}