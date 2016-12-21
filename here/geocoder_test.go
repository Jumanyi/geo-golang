package here_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/here"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var appID = os.Getenv("HERE_APP_ID")
var appCode = os.Getenv("HERE_APP_CODE")

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := here.Geocoder(appID, appCode, 100, ts.URL+"/")
	location, err := geocoder.Geocode("60 Collins St, Melbourne VIC 3000")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.81375, Lng: 144.97176}, location)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := here.Geocoder(appID, appCode, 100, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(-37.81375, 144.97176)
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(address, "56-64 Collins St"))
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := here.Geocoder(appID, appCode, 100, ts.URL+"/")
	_, err := geocoder.ReverseGeocode(-37.81375, 164.97176)
	assert.Equal(t, err, geo.ErrNoResult)
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

const (
	response1 = `{
   "Response":{
      "MetaInfo":{
         "Timestamp":"2016-05-13T07:57:11.878+0000"
      },
      "View":[
         {
            "_type":"SearchResultsViewType",
            "ViewId":0,
            "Result":[
               {
                  "Relevance":1.0,
                  "MatchLevel":"houseNumber",
                  "MatchQuality":{
                     "State":1.0,
                     "City":1.0,
                     "Street":[
                        1.0
                     ],
                     "HouseNumber":1.0,
                     "PostalCode":1.0
                  },
                  "MatchType":"pointAddress",
                  "Location":{
                     "LocationId":"NT_o.OOjClAWfILi-O4OxzVCA_2AD",
                     "LocationType":"address",
                     "DisplayPosition":{
                        "Latitude":-37.81375,
                        "Longitude":144.97176
                     },
                     "NavigationPosition":[
                        {
                           "Latitude":-37.81393,
                           "Longitude":144.97185
                        }
                     ],
                     "MapView":{
                        "TopLeft":{
                           "Latitude":-37.8126258,
                           "Longitude":144.970337
                        },
                        "BottomRight":{
                           "Latitude":-37.8148742,
                           "Longitude":144.973183
                        }
                     },
                     "Address":{
                        "Label":"60 Collins St, Melbourne VIC 3000, Australia",
                        "Country":"AUS",
                        "State":"VIC",
                        "City":"Melbourne",
                        "Street":"Collins St",
                        "HouseNumber":"60",
                        "PostalCode":"3000",
                        "AdditionalData":[
                           {
                              "value":"Australia",
                              "key":"CountryName"
                           },
                           {
                              "value":"Victoria",
                              "key":"StateName"
                           }
                        ]
                     }
                  }
               }
            ]
         }
      ]
   }
}`
	response2 = `{
   "Response":{
      "MetaInfo":{
         "Timestamp":"2016-05-13T07:57:42.917+0000",
         "NextPageInformation":"2"
      },
      "View":[
         {
            "_type":"SearchResultsViewType",
            "ViewId":0,
            "Result":[
               {
                  "Relevance":1.0,
                  "Distance":0.0,
                  "MatchLevel":"houseNumber",
                  "MatchQuality":{
                     "Country":1.0,
                     "State":1.0,
                     "City":1.0,
                     "Street":[
                        1.0
                     ],
                     "HouseNumber":1.0,
                     "PostalCode":1.0
                  },
                  "MatchType":"pointAddress",
                  "Location":{
                     "LocationId":"NT_o.OOjClAWfILi-O4OxzVCA_1YTL2QD",
                     "LocationType":"address",
                     "DisplayPosition":{
                        "Latitude":-37.81375,
                        "Longitude":144.97176
                     },
                     "NavigationPosition":[
                        {
                           "Latitude":-37.81392,
                           "Longitude":144.97184
                        }
                     ],
                     "MapView":{
                        "TopLeft":{
                           "Latitude":-37.8126258,
                           "Longitude":144.970337
                        },
                        "BottomRight":{
                           "Latitude":-37.8148742,
                           "Longitude":144.973183
                        }
                     },
                     "Address":{
                        "Label":"56-64 Collins St, Melbourne VIC 3000, Australia",
                        "Country":"AUS",
                        "State":"VIC",
                        "City":"Melbourne",
                        "Street":"Collins St",
                        "HouseNumber":"56-64",
                        "PostalCode":"3000",
                        "AdditionalData":[
                           {
                              "value":"Australia",
                              "key":"CountryName"
                           },
                           {
                              "value":"Victoria",
                              "key":"StateName"
                           }
                        ]
                     },
                     "MapReference":{
                        "ReferenceId":"718734408",
                        "MapId":"NXAM16108",
                        "MapVersion":"Q1/2016",
                        "MapReleaseDate":"2016-05-04",
                        "Spot":0.43,
                        "SideOfStreet":"left",
                        "CountryId":"1469256839",
                        "StateId":"1469285956",
                        "CityId":"1469261839",
                        "AddressId":"335909793"
                     }
                  }
               },
               {
                  "Relevance":1.0,
                  "Distance":13.9,
                  "MatchLevel":"houseNumber",
                  "MatchQuality":{
                     "Country":1.0,
                     "State":1.0,
                     "City":1.0,
                     "Street":[
                        1.0
                     ],
                     "HouseNumber":1.0,
                     "PostalCode":1.0
                  },
                  "MatchType":"pointAddress",
                  "Location":{
                     "LocationId":"NT_o.OOjClAWfILi-O4OxzVCA_1ITL1QD",
                     "LocationType":"address",
                     "DisplayPosition":{
                        "Latitude":-37.81371,
                        "Longitude":144.97191
                     },
                     "NavigationPosition":[
                        {
                           "Latitude":-37.81388,
                           "Longitude":144.97198
                        }
                     ],
                     "MapView":{
                        "TopLeft":{
                           "Latitude":-37.8125858,
                           "Longitude":144.970487
                        },
                        "BottomRight":{
                           "Latitude":-37.8148342,
                           "Longitude":144.973333
                        }
                     },
                     "Address":{
                        "Label":"52-54 Collins St, Melbourne VIC 3000, Australia",
                        "Country":"AUS",
                        "State":"VIC",
                        "City":"Melbourne",
                        "Street":"Collins St",
                        "HouseNumber":"52-54",
                        "PostalCode":"3000",
                        "AdditionalData":[
                           {
                              "value":"Australia",
                              "key":"CountryName"
                           },
                           {
                              "value":"Victoria",
                              "key":"StateName"
                           }
                        ]
                     },
                     "MapReference":{
                        "ReferenceId":"1152786853",
                        "MapId":"NXAM16108",
                        "MapVersion":"Q1/2016",
                        "MapReleaseDate":"2016-05-04",
                        "Spot":0.28,
                        "SideOfStreet":"left",
                        "CountryId":"1469256839",
                        "StateId":"1469285956",
                        "CityId":"1469261839",
                        "AddressId":"335008017"
                     }
                  }
               },
               {
                  "Relevance":1.0,
                  "Distance":20.4,
                  "MatchLevel":"houseNumber",
                  "MatchQuality":{
                     "Country":1.0,
                     "State":1.0,
                     "City":1.0,
                     "Street":[
                        1.0
                     ],
                     "HouseNumber":1.0,
                     "PostalCode":1.0
                  },
                  "MatchType":"interpolated",
                  "Location":{
                     "LocationId":"NT_o.OOjClAWfILi-O4OxzVCA_l_718734407_L_2ID",
                     "LocationType":"address",
                     "DisplayPosition":{
                        "Latitude":-37.8138635,
                        "Longitude":144.9715768
                     },
                     "NavigationPosition":[
                        {
                           "Latitude":-37.813995,
                           "Longitude":144.971615
                        }
                     ],
                     "MapView":{
                        "TopLeft":{
                           "Latitude":-37.8127394,
                           "Longitude":144.9701539
                        },
                        "BottomRight":{
                           "Latitude":-37.8149877,
                           "Longitude":144.9729998
                        }
                     },
                     "Address":{
                        "Label":"62 Collins St, Melbourne VIC 3000, Australia",
                        "Country":"AUS",
                        "State":"VIC",
                        "City":"Melbourne",
                        "Street":"Collins St",
                        "HouseNumber":"62",
                        "PostalCode":"3000",
                        "AdditionalData":[
                           {
                              "value":"Australia",
                              "key":"CountryName"
                           },
                           {
                              "value":"Victoria",
                              "key":"StateName"
                           }
                        ]
                     },
                     "MapReference":{
                        "ReferenceId":"718734407",
                        "MapId":"NXAM16108",
                        "MapVersion":"Q1/2016",
                        "MapReleaseDate":"2016-05-04",
                        "Spot":0.5,
                        "SideOfStreet":"left",
                        "CountryId":"1469256839",
                        "StateId":"1469285956",
                        "CityId":"1469261839"
                     }
                  }
               },
               {
                  "Relevance":1.0,
                  "Distance":21.7,
                  "MatchLevel":"houseNumber",
                  "MatchQuality":{
                     "Country":1.0,
                     "State":1.0,
                     "City":1.0,
                     "Street":[
                        1.0
                     ],
                     "HouseNumber":1.0,
                     "PostalCode":1.0
                  },
                  "MatchType":"interpolated",
                  "Location":{
                     "LocationId":"NT_5oLXk7SzYlmd8YgeTgyzVB_l_718734413_R_4ID",
                     "LocationType":"address",
                     "DisplayPosition":{
                        "Latitude":-37.8138589,
                        "Longitude":144.9715545
                     },
                     "NavigationPosition":[
                        {
                           "Latitude":-37.8139275,
                           "Longitude":144.9714075
                        }
                     ],
                     "MapView":{
                        "TopLeft":{
                           "Latitude":-37.8127347,
                           "Longitude":144.9701316
                        },
                        "BottomRight":{
                           "Latitude":-37.814983,
                           "Longitude":144.9729775
                        }
                     },
                     "Address":{
                        "Label":"82 Exhibition St, Melbourne VIC 3000, Australia",
                        "Country":"AUS",
                        "State":"VIC",
                        "City":"Melbourne",
                        "Street":"Exhibition St",
                        "HouseNumber":"82",
                        "PostalCode":"3000",
                        "AdditionalData":[
                           {
                              "value":"Australia",
                              "key":"CountryName"
                           },
                           {
                              "value":"Victoria",
                              "key":"StateName"
                           }
                        ]
                     },
                     "MapReference":{
                        "ReferenceId":"718734413",
                        "MapId":"NXAM16108",
                        "MapVersion":"Q1/2016",
                        "MapReleaseDate":"2016-05-04",
                        "Spot":0.75,
                        "SideOfStreet":"right",
                        "CountryId":"1469256839",
                        "StateId":"1469285956",
                        "CityId":"1469261839"
                     }
                  }
               },
               {
                  "Relevance":1.0,
                  "Distance":25.7,
                  "MatchLevel":"houseNumber",
                  "MatchQuality":{
                     "Country":1.0,
                     "State":1.0,
                     "City":1.0,
                     "Street":[
                        1.0
                     ],
                     "HouseNumber":1.0,
                     "PostalCode":1.0
                  },
                  "MatchType":"interpolated",
                  "Location":{
                     "LocationId":"NT_5oLXk7SzYlmd8YgeTgyzVB_l_718734414_R_4QD",
                     "LocationType":"address",
                     "DisplayPosition":{
                        "Latitude":-37.8136944,
                        "Longitude":144.9714755
                     },
                     "NavigationPosition":[
                        {
                           "Latitude":-37.8137639,
                           "Longitude":144.9713292
                        }
                     ],
                     "MapView":{
                        "TopLeft":{
                           "Latitude":-37.8125702,
                           "Longitude":144.9700526
                        },
                        "BottomRight":{
                           "Latitude":-37.8148185,
                           "Longitude":144.9728985
                        }
                     },
                     "Address":{
                        "Label":"84 Exhibition St, Melbourne VIC 3000, Australia",
                        "Country":"AUS",
                        "State":"VIC",
                        "City":"Melbourne",
                        "Street":"Exhibition St",
                        "HouseNumber":"84",
                        "PostalCode":"3000",
                        "AdditionalData":[
                           {
                              "value":"Australia",
                              "key":"CountryName"
                           },
                           {
                              "value":"Victoria",
                              "key":"StateName"
                           }
                        ]
                     },
                     "MapReference":{
                        "ReferenceId":"718734414",
                        "MapId":"NXAM16108",
                        "MapVersion":"Q1/2016",
                        "MapReleaseDate":"2016-05-04",
                        "Spot":0.47,
                        "SideOfStreet":"right",
                        "CountryId":"1469256839",
                        "StateId":"1469285956",
                        "CityId":"1469261839"
                     }
                  }
               },
               {
                  "Relevance":1.0,
                  "Distance":25.8,
                  "MatchLevel":"houseNumber",
                  "MatchQuality":{
                     "Country":1.0,
                     "State":1.0,
                     "City":1.0,
                     "Street":[
                        1.0
                     ],
                     "HouseNumber":1.0,
                     "PostalCode":1.0
                  },
                  "MatchType":"interpolated",
                  "Location":{
                     "LocationId":"NT_o.OOjClAWfILi-O4OxzVCA_l_1152786854_L_0gD",
                     "LocationType":"address",
                     "DisplayPosition":{
                        "Latitude":-37.8137284,
                        "Longitude":144.9720524
                     },
                     "NavigationPosition":[
                        {
                           "Latitude":-37.81386,
                           "Longitude":144.97209
                        }
                     ],
                     "MapView":{
                        "TopLeft":{
                           "Latitude":-37.8126043,
                           "Longitude":144.9706295
                        },
                        "BottomRight":{
                           "Latitude":-37.8148526,
                           "Longitude":144.9734754
                        }
                     },
                     "Address":{
                        "Label":"48 Collins St, Melbourne VIC 3000, Australia",
                        "Country":"AUS",
                        "State":"VIC",
                        "City":"Melbourne",
                        "Street":"Collins St",
                        "HouseNumber":"48",
                        "PostalCode":"3000",
                        "AdditionalData":[
                           {
                              "value":"Australia",
                              "key":"CountryName"
                           },
                           {
                              "value":"Victoria",
                              "key":"StateName"
                           }
                        ]
                     },
                     "MapReference":{
                        "ReferenceId":"1152786854",
                        "MapId":"NXAM16108",
                        "MapVersion":"Q1/2016",
                        "MapReleaseDate":"2016-05-04",
                        "Spot":0.0,
                        "SideOfStreet":"left",
                        "CountryId":"1469256839",
                        "StateId":"1469285956",
                        "CityId":"1469261839"
                     }
                  }
               },
               {
                  "Relevance":1.0,
                  "Distance":29.7,
                  "MatchLevel":"houseNumber",
                  "MatchQuality":{
                     "Country":1.0,
                     "State":1.0,
                     "City":1.0,
                     "Street":[
                        1.0
                     ],
                     "HouseNumber":1.0,
                     "PostalCode":1.0
                  },
                  "MatchType":"interpolated",
                  "Location":{
                     "LocationId":"NT_5oLXk7SzYlmd8YgeTgyzVB_l_718734405_R_3QD",
                     "LocationType":"address",
                     "DisplayPosition":{
                        "Latitude":-37.8139922,
                        "Longitude":144.9716167
                     },
                     "NavigationPosition":[
                        {
                           "Latitude":-37.8140612,
                           "Longitude":144.97147
                        }
                     ],
                     "MapView":{
                        "TopLeft":{
                           "Latitude":-37.8128681,
                           "Longitude":144.9701937
                        },
                        "BottomRight":{
                           "Latitude":-37.8151164,
                           "Longitude":144.9730397
                        }
                     },
                     "Address":{
                        "Label":"74 Exhibition St, Melbourne VIC 3000, Australia",
                        "Country":"AUS",
                        "State":"VIC",
                        "City":"Melbourne",
                        "Street":"Exhibition St",
                        "HouseNumber":"74",
                        "PostalCode":"3000",
                        "AdditionalData":[
                           {
                              "value":"Australia",
                              "key":"CountryName"
                           },
                           {
                              "value":"Victoria",
                              "key":"StateName"
                           }
                        ]
                     },
                     "MapReference":{
                        "ReferenceId":"718734405",
                        "MapId":"NXAM16108",
                        "MapVersion":"Q1/2016",
                        "MapReleaseDate":"2016-05-04",
                        "Spot":0.88,
                        "SideOfStreet":"right",
                        "CountryId":"1469256839",
                        "StateId":"1469285956",
                        "CityId":"1469261839"
                     }
                  }
               },
               {
                  "Relevance":1.0,
                  "Distance":42.3,
                  "MatchLevel":"street",
                  "MatchQuality":{
                     "Country":1.0,
                     "State":1.0,
                     "City":1.0,
                     "Street":[
                        1.0
                     ],
                     "PostalCode":1.0
                  },
                  "Location":{
                     "LocationId":"NT_Qs4XgGPRnjfYuvDCd6nxHB_l_134268791_L",
                     "LocationType":"address",
                     "DisplayPosition":{
                        "Latitude":-37.81338,
                        "Longitude":144.97165
                     },
                     "NavigationPosition":[
                        {
                           "Latitude":-37.81338,
                           "Longitude":144.97165
                        }
                     ],
                     "MapView":{
                        "TopLeft":{
                           "Latitude":-37.81293,
                           "Longitude":144.97144
                        },
                        "BottomRight":{
                           "Latitude":-37.81338,
                           "Longitude":144.97165
                        }
                     },
                     "Address":{
                        "Label":"McGraths Ln, Melbourne VIC 3000, Australia",
                        "Country":"AUS",
                        "State":"VIC",
                        "City":"Melbourne",
                        "Street":"McGraths Ln",
                        "PostalCode":"3000",
                        "AdditionalData":[
                           {
                              "value":"Australia",
                              "key":"CountryName"
                           },
                           {
                              "value":"Victoria",
                              "key":"StateName"
                           }
                        ]
                     },
                     "MapReference":{
                        "ReferenceId":"134268791",
                        "MapId":"NXAM16108",
                        "MapVersion":"Q1/2016",
                        "MapReleaseDate":"2016-05-04",
                        "Spot":0.0,
                        "SideOfStreet":"left",
                        "CountryId":"1469256839",
                        "StateId":"1469285956",
                        "CityId":"1469261839"
                     }
                  }
               },
               {
                  "Relevance":1.0,
                  "Distance":42.5,
                  "MatchLevel":"houseNumber",
                  "MatchQuality":{
                     "Country":1.0,
                     "State":1.0,
                     "City":1.0,
                     "Street":[
                        1.0
                     ],
                     "HouseNumber":1.0,
                     "PostalCode":1.0
                  },
                  "MatchType":"pointAddress",
                  "Location":{
                     "LocationId":"NT_o.OOjClAWfILi-O4OxzVCA_zYTL1AD",
                     "LocationType":"address",
                     "DisplayPosition":{
                        "Latitude":-37.81361,
                        "Longitude":144.97221
                     },
                     "NavigationPosition":[
                        {
                           "Latitude":-37.81379,
                           "Longitude":144.9723
                        }
                     ],
                     "MapView":{
                        "TopLeft":{
                           "Latitude":-37.8124858,
                           "Longitude":144.970787
                        },
                        "BottomRight":{
                           "Latitude":-37.8147342,
                           "Longitude":144.973633
                        }
                     },
                     "Address":{
                        "Label":"36-50 Collins St, Melbourne VIC 3000, Australia",
                        "Country":"AUS",
                        "State":"VIC",
                        "City":"Melbourne",
                        "Street":"Collins St",
                        "HouseNumber":"36-50",
                        "PostalCode":"3000",
                        "AdditionalData":[
                           {
                              "value":"Australia",
                              "key":"CountryName"
                           },
                           {
                              "value":"Victoria",
                              "key":"StateName"
                           }
                        ]
                     },
                     "MapReference":{
                        "ReferenceId":"1152786855",
                        "MapId":"NXAM16108",
                        "MapVersion":"Q1/2016",
                        "MapReleaseDate":"2016-05-04",
                        "Spot":0.28,
                        "SideOfStreet":"left",
                        "CountryId":"1469256839",
                        "StateId":"1469285956",
                        "CityId":"1469261839",
                        "AddressId":"335008015"
                     }
                  }
               },
               {
                  "Relevance":1.0,
                  "Distance":46.2,
                  "MatchLevel":"houseNumber",
                  "MatchQuality":{
                     "Country":1.0,
                     "State":1.0,
                     "City":1.0,
                     "Street":[
                        1.0
                     ],
                     "HouseNumber":1.0,
                     "PostalCode":1.0
                  },
                  "MatchType":"interpolated",
                  "Location":{
                     "LocationId":"NT_5oLXk7SzYlmd8YgeTgyzVB_l_718734404_R_1gD",
                     "LocationType":"address",
                     "DisplayPosition":{
                        "Latitude":-37.8141628,
                        "Longitude":144.9716968
                     },
                     "NavigationPosition":[
                        {
                           "Latitude":-37.814235,
                           "Longitude":144.9715525
                        }
                     ],
                     "MapView":{
                        "TopLeft":{
                           "Latitude":-37.8130387,
                           "Longitude":144.9702738
                        },
                        "BottomRight":{
                           "Latitude":-37.815287,
                           "Longitude":144.9731198
                        }
                     },
                     "Address":{
                        "Label":"58 Exhibition St, Melbourne VIC 3000, Australia",
                        "Country":"AUS",
                        "State":"VIC",
                        "City":"Melbourne",
                        "Street":"Exhibition St",
                        "HouseNumber":"58",
                        "PostalCode":"3000",
                        "AdditionalData":[
                           {
                              "value":"Australia",
                              "key":"CountryName"
                           },
                           {
                              "value":"Victoria",
                              "key":"StateName"
                           }
                        ]
                     },
                     "MapReference":{
                        "ReferenceId":"718734404",
                        "MapId":"NXAM16108",
                        "MapVersion":"Q1/2016",
                        "MapReleaseDate":"2016-05-04",
                        "Spot":0.88,
                        "SideOfStreet":"right",
                        "CountryId":"1469256839",
                        "StateId":"1469285956",
                        "CityId":"1469261839"
                     }
                  }
               }
            ]
         }
      ]
   }
}`
	response3 = `{
   "Response":{
      "MetaInfo":{
         "Timestamp":"2016-05-13T07:57:43.402+0000"
      },
      "View":[

      ]
   }
}`
)
