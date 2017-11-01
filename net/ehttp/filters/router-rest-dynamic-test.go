//Copyright 2017 SabzCity
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package filters

var dynaimcRouteTest = map[string][]DynamicRestRoute{
	"apis.sabz.city": []DynamicRestRoute{
		DynamicRestRoute{
			Path:        "/auth/v0/",
			Type:        "Container",
			Destination: "v0.auth.sabzcity.containers.sabz.city:50000"},
		DynamicRestRoute{
			Path:        "/usersinfo/v0/",
			Type:        "Container",
			Destination: "v0.usersinfo.sabzcity.containers.sabz.city:50005"},
		DynamicRestRoute{
			Path:        "/groups/v0/",
			Type:        "Container",
			Destination: "v0.groups.sabzcity.containers.sabz.city:50010"},
		DynamicRestRoute{
			Path:        "/storage/v0/",
			Type:        "Container",
			Destination: "v0.storage.sabzcity.containers.sabz.city:50015"},
		DynamicRestRoute{
			Path:        "/kvs/v0/",
			Type:        "Container",
			Destination: "v0.kvs.sabzcity.containers.sabz.city:50020"},
		DynamicRestRoute{
			Path:        "/logs/v0/",
			Type:        "Container",
			Destination: "v0.logs.sabzcity.containers.sabz.city:50025"},
		DynamicRestRoute{
			Path:        "/domains/v0/",
			Type:        "Container",
			Destination: "v0.domains.sabzcity.containers.sabz.city:50030"},
		DynamicRestRoute{
			Path:        "/products/v0/",
			Type:        "Container",
			Destination: "v0.products.sabzcity.containers.sabz.city:50035"},
		DynamicRestRoute{
			Path:        "/financials/v0/",
			Type:        "Container",
			Destination: "v0.financials.sabzcity.containers.sabz.city:50050"},
		DynamicRestRoute{
			Path:        "/complexes/v0/",
			Type:        "Container",
			Destination: "v0.complexes.sabzcity.containers.sabz.city:50060"},
		DynamicRestRoute{
			Path:        "/coordinates/v0/",
			Type:        "Container",
			Destination: "v0.coordinates.sabzcity.containers.sabz.city:50065"},
		DynamicRestRoute{
			Path:        "/units/v0/",
			Type:        "Container",
			Destination: "v0.units.sabzcity.containers.sabz.city:50070"},
		DynamicRestRoute{
			Path:        "/wiki/v0/",
			Type:        "Container",
			Destination: "v0.wiki.sabzcity.containers.sabz.city:50100"},
		DynamicRestRoute{
			Path:        "/invoices/v0/",
			Type:        "Container",
			Destination: "v0.invoices.sabzcity.containers.sabz.city:50105"},
		// TEST
		DynamicRestRoute{
			Path:        "/example/",
			Destination: "http://www.example.com",
			Type:        "Proxy"}},
	"sabz.city": []DynamicRestRoute{
		DynamicRestRoute{
			Path:        "*",
			Destination: "http://sabz.city:8000",
			Type:        "Proxy"}},
	"www.sabz.city": []DynamicRestRoute{
		DynamicRestRoute{
			Path:        "*",
			Destination: "http://www.sabz.city:8000",
			Type:        "Proxy"}},
	"shop.sabz.city": []DynamicRestRoute{
		DynamicRestRoute{
			Path:        "*",
			Destination: "http://shop.sabz.city:8000",
			Type:        "Proxy"}},
	"myorg.sabz.city": []DynamicRestRoute{
		DynamicRestRoute{
			Path:        "*",
			Destination: "http://myorg.sabz.city:8000",
			Type:        "Proxy"}},
	"accounting.sabz.city": []DynamicRestRoute{
		DynamicRestRoute{
			Path:        "*",
			Destination: "http://accounting.sabz.city:8000/",
			Type:        "Proxy"}},
	"transport.sabz.city": []DynamicRestRoute{
		DynamicRestRoute{
			Path:        "*",
			Destination: "http://transport.sabz.city:8000",
			Type:        "Proxy"}},
	"my.sabz.city": []DynamicRestRoute{
		DynamicRestRoute{
			Path:        "*",
			Destination: "http://my.sabz.city:8000",
			Type:        "Proxy"}}}
