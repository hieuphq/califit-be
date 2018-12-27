package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Token token response
var Token = MediaType("application/vnd.token+json", func() {
	Description("A token")
	Attributes(func() {
		Attribute("token", String, "A JWT token")
	})

	View("default", func() {
		Attribute("token")
	})
})

// LoginRes login response
var LoginRes = MediaType("application/vnd.login_response+json", func() {
	Description("Login response")

	Attribute("token", String)
	Attribute("user_info", UserInfo)

	View("default", func() {
		Attribute("token")
		Attribute("user_info")
	})
})

// UserInfo user response
var UserInfo = MediaType("application/vnd.user_info+json", func() {
	Description("User info")

	Attribute("name", String)
	Attribute("email", String)
	Attribute("phone", String)
	Attribute("name", String)
	Attribute("company_address", String)

	View("default", func() {
		Attribute("name")
		Attribute("email")
		Attribute("phone")
		Attribute("name")
		Attribute("company_address")
	})
})

// City response data
var City = MediaType("application/vnd.city+json", func() {
	Description("a city detail")

	Attribute("id", String)
	Attribute("name", String)
	Attribute("created_at", DateTime)

	View("default", func() {
		Attribute("id", String)
		Attribute("name", String)
		Attribute("created_at", DateTime)
	})

	View("general", func() {
		Attribute("id", String)
		Attribute("name", String)
		Attribute("created_at", DateTime)
	})
})

// CityList enum of city
var CityList = MediaType("application/vnd.cities+json", func() {
	Description("A list of City")

	Attribute("data", ArrayOf(City, func() {
		View("general")
	}))
	Attribute("paginate", Paginate)

	View("default", func() {
		Attribute("data")
		Attribute("paginate")
	})
})

// Address response data
var Address = MediaType("application/vnd.address+json", func() {
	Description("a address detail")

	Attribute("id", String)
	Attribute("name", String)
	Attribute("street1", String)
	Attribute("street2", String)
	Attribute("city", String)
	Attribute("state", String)
	Attribute("country", String)
	Attribute("zip_code", String)
	Attribute("lat", Number)
	Attribute("lng", Number)
	Attribute("created_at", DateTime)

	View("default", func() {
		Attribute("id", String)
		Attribute("name", String)
		Attribute("street1", String)
		Attribute("street2", String)
		Attribute("city", String)
		Attribute("state", String)
		Attribute("country", String)
		Attribute("zip_code", String)
		Attribute("lat", Number)
		Attribute("lng", Number)
		Attribute("created_at", DateTime)
	})

	View("general", func() {
		Attribute("id", String)
		Attribute("name", String)
		Attribute("street1", String)
		Attribute("street2", String)
		Attribute("city", String)
		Attribute("state", String)
		Attribute("country", String)
		Attribute("zip_code", String)
		Attribute("lat", Number)
		Attribute("lng", Number)
		Attribute("created_at", DateTime)
	})
})

// Center response data
var Center = MediaType("application/vnd.center+json", func() {
	Description("a center detail")

	Attribute("id", String)
	Attribute("name", String)
	Attribute("created_at", DateTime)
	Attribute("address", Address)
	Attribute("city", City)
	Attribute("address_id", String)
	Attribute("city_id", String)

	View("default", func() {
		Attribute("id", String)
		Attribute("name", String)
		Attribute("created_at", DateTime)
		Attribute("address", Address)
		Attribute("city", City)
		Attribute("address_id", String)
		Attribute("city_id", String)
	})

	View("general", func() {
		Attribute("id", String)
		Attribute("name", String)
		Attribute("created_at", DateTime)
		Attribute("address", Address)
		Attribute("city", City)
		Attribute("address_id", String)
		Attribute("city_id", String)
	})
})

// CenterList enum of center
var CenterList = MediaType("application/vnd.centers+json", func() {
	Description("A list of Center")

	Attribute("data", ArrayOf(Center, func() {
		View("general")
	}))
	Attribute("paginate", Paginate)

	View("default", func() {
		Attribute("data")
		Attribute("paginate")
	})
})

// Paginate response
var Paginate = MediaType("application/vnd.paginate+json", func() {
	Description("pagination of a resources")
	Attribute("current_page", Integer)
	Attribute("total_page", Integer)
	Attribute("total_item", Integer)

	View("default", func() {
		Attribute("current_page")
		Attribute("total_page")
		Attribute("total_item")
	})
})

// Schedule response data
var Schedule = MediaType("application/vnd.schedule+json", func() {
	Description("a schedule detail")

	Attribute("id", String)
	Attribute("center_id", String)
	Attribute("center", Center)
	Attribute("start_at", DateTime)
	Attribute("end_at", DateTime)

	View("default", func() {
		Attribute("id", String)
		Attribute("center_id", String)
		Attribute("center", Center)
		Attribute("start_at", DateTime)
		Attribute("end_at", DateTime)
	})

	View("general", func() {
		Attribute("id", String)
		Attribute("center_id", String)
		Attribute("center", Center)
		Attribute("start_at", DateTime)
		Attribute("end_at", DateTime)
	})
})

// ScheduleList enum of schedule list by center ID
var ScheduleList = MediaType("application/vnd.schedules+json", func() {
	Description("A list of Schedule by Center ID")

	Attribute("data", ArrayOf(Schedule, func() {
		View("general")
	}))
	Attribute("paginate", Paginate)

	View("default", func() {
		Attribute("data")
		Attribute("paginate")
	})
})
