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

import (
	"github.com/SabzCity/go-library/encoding/ejson"
	"github.com/SabzCity/go-library/net/edns/ednsutil"
)

func init() {
	// Test
	var sabzcity ednsutil.DNS
	ejson.Unmarshal(sabzCityZone, &sabzcity)
	StaticZoneCache.Set(&sabzcity)

	var ael ednsutil.DNS
	ejson.Unmarshal(aelIrZone, &ael)
	StaticZoneCache.Set(&ael)

	var sabzcitynet ednsutil.DNS
	ejson.Unmarshal(sabzcityNetZone, &sabzcitynet)
	StaticZoneCache.Set(&sabzcitynet)

	var shahr ednsutil.DNS
	ejson.Unmarshal(shahrsabzIrZone, &shahr)
	StaticZoneCache.Set(&shahr)

	var szc ednsutil.DNS
	ejson.Unmarshal(szcIrZone, &szc)
	StaticZoneCache.Set(&szc)

	var sss ednsutil.DNS
	ejson.Unmarshal(xnngbrecd8iComZone, &sss)
	StaticZoneCache.Set(&sss)
}

var sabzCityZone = `
{
    "Origin": "sabz.city.",
    "TTL": 3600,
    "IN": {
        "sabz.city.": {
            "SOA": {
                "MNAME": "ns1.sabz.city.",
                "RNAME": "it.sabz.city.",
                "SERIAL": 3600,
                "REFRESH": 2017071600,
                "RETRY": 3600,
                "EXPIRE": 3600,
                "MINIMUM": 2419200
            },
            "NSDNAME": [
                "ns1.sabz.city.",
                "ns2.sabz.city.",
                "ns3.sabz.city.",
                "ns4.sabz.city.",
                "ns5.sabz.city."
            ],
            "A": [
                "46.4.162.78"
            ],
            "MX": [{
                "PREFERENCE": 0,
                "EXCHANGE": "mail.sabz.city."
            }],
            "TXTDATA": [
                "google-site-verification=btDdt18Sr4U9HsiVsfmZkIoGsZhXgIHOuvO3Hn3cvUY"
            ]
        },
        "ns1.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "ns2.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "ns3.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "ns4.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "apis.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "containers.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "my.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "myorg.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "accounting.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "blog.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "wiki.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "shop.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "transport.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "services.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "need.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "health.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "maps.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "life.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "drive.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "communication.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "music.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "video.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "sg.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "developers.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "www.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "code.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "smtp.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "ftp.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "mail.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        },
        "pop.sabz.city.": {
            "A": [
                "46.4.162.78"
            ]
        }
    }
}
`

var aelIrZone = `
{
    "Origin": "ael.ir.",
    "TTL": 3600,
    "IN": {
        "ael.ir.": {
            "SOA": {
                "MName": "ns1.sabz.city.",
                "RName": "it.sabz.city.",
                "Serial": 3600,
                "Refresh": 2017071600,
                "Retry": 3600,
                "Expire": 3600,
                "Minimum": 2419200
            },
            "NS": [
                "ns1.sabz.city.",
                "ns2.sabz.city.",
                "ns3.sabz.city.",
                "ns4.sabz.city.",
                "ns5.sabz.city."
            ],
            "A": [
                "46.4.162.78"
            ],
            "MX": [{
                "Preference": 10,
                "Host": "mail.sabz.city."
            }],
            "TXT": [
                "google-site-verification=nOA_w_cvyMsDgfTzZaED41gcDukL1EauQ5aB0gFM8iI"
            ]
        },
        "www.ael.ir.": {
            "A": [
                "46.4.162.78"
            ]
        }
    }
}
`

var sabzcityNetZone = `
{
    "Origin": "sabzcity.net.",
    "TTL": 3600,
    "IN": {
        "sabzcity.net.": {
            "SOA": {
                "MName": "ns1.sabz.city.",
                "RName": "it.sabz.city.",
                "Serial": 3600,
                "Refresh": 2017071600,
                "Retry": 3600,
                "Expire": 3600,
                "Minimum": 2419200
            },
            "NS": [
                "ns1.sabz.city.",
                "ns2.sabz.city.",
                "ns3.sabz.city.",
                "ns4.sabz.city.",
                "ns5.sabz.city."
            ],
            "A": [
                "46.4.162.78"
            ],
            "MX": [{
                "Preference": 10,
                "Host": "mail.sabz.city."
            }],
            "TXT": [
                "google-site-verification=h2p2tn0Yk_xf0gG-s468APN8sGKYFeM2yygWBLaDMBw"
            ]
        },
        "www.sabzcity.net.": {
            "A": [
                "46.4.162.78"
            ]
        }
    }
}
`

var shahrsabzIrZone = `
{
    "Origin": "shahrsabz.ir.",
    "TTL": 3600,
    "IN": {
        "shahrsabz.ir.": {
            "SOA": {
                "MName": "ns1.sabz.city.",
                "RName": "it.sabz.city.",
                "Serial": 3600,
                "Refresh": 2017071600,
                "Retry": 3600,
                "Expire": 3600,
                "Minimum": 2419200
            },
            "NS": [
                "ns1.sabz.city.",
                "ns2.sabz.city.",
                "ns3.sabz.city.",
                "ns4.sabz.city.",
                "ns5.sabz.city."
            ],
            "A": [
                "46.4.162.78"
            ],
            "MX": [{
                "Preference": 10,
                "Host": "mail.sabz.city."
            }],
            "TXT": [
                "google-site-verification=nOZSivdxaU1VPR3As46nLh-bitoI-JnrI8On65o8pxM"
            ]
        },
        "www.shahrsabz.ir.": {
            "A": [
                "46.4.162.78"
            ]
        }
    }
}
`

var szcIrZone = `
{
    "Origin": "szc.ir.",
    "TTL": 3600,
    "IN": {
        "szc.ir.": {
            "SOA": {
                "MName": "ns1.sabz.city.",
                "RName": "it.sabz.city.",
                "Serial": 3600,
                "Refresh": 2017071600,
                "Retry": 3600,
                "Expire": 3600,
                "Minimum": 2419200
            },
            "NS": [
                "ns1.sabz.city.",
                "ns2.sabz.city.",
                "ns3.sabz.city.",
                "ns4.sabz.city.",
                "ns5.sabz.city."
            ],
            "A": [
                "46.4.162.78"
            ],
            "MX": [{
                "Preference": 10,
                "Host": "mail.sabz.city."
            }],
            "TXT": [
                "google-site-verification=GmUg2bc4Vct729YhD50ukCoxZ4H0VibG9YlNmylAsXQ"
            ]
        },
        "www.szc.ir.": {
            "A": [
                "46.4.162.78"
            ]
        }
    }
}
`

var xnngbrecd8iComZone = `
{
    "Origin": "xn--ngbrecd8i.com.",
    "TTL": 3600,
    "IN": {
        "xn--ngbrecd8i.com.": {
            "SOA": {
                "MName": "ns1.sabz.city.",
                "RName": "it.sabz.city.",
                "Serial": 3600,
                "Refresh": 2017071600,
                "Retry": 3600,
                "Expire": 3600,
                "Minimum": 2419200
            },
            "NS": [
                "ns1.sabz.city.",
                "ns2.sabz.city.",
                "ns3.sabz.city.",
                "ns4.sabz.city.",
                "ns5.sabz.city."
            ],
            "A": [
                "46.4.162.78"
            ],
            "MX": [{
                "Preference": 10,
                "Host": "mail.sabz.city."
            }],
            "TXT": [
                "google-site-verification=yf9NNx4VZ9oWCppegulm72JXMmVjpuVSJg6ru9JN5x8"
            ]
        },
        "www.xn--ngbrecd8i.com.": {
            "A": [
                "46.4.162.78"
            ]
        }
    }
}
`
