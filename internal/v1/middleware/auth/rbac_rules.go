package auth

// "WkCTJ7F4xI" == DinDin-Admin
// "DbtYHGhUW7" == DinDin-User
const _CASBIN_RULES = `
[
  {
    "ptype": "p",
    "v0": "WkCTJ7F4xI",
    "v1": "/hamadan/v1.1/*",
    "v2": "GET"
  },
  {
    "ptype": "p",
    "v0": "WkCTJ7F4xI",
    "v1": "/hamadan/v1.1/*",
    "v2": "PUT"
  },
  {
    
    "ptype": "p",
    "v0": "WkCTJ7F4xI",
    "v1": "/hamadan/v1.1/*",
    "v2": "PATCH"
  },
  {
    
    "ptype": "p",
    "v0": "WkCTJ7F4xI",
    "v1": "/hamadan/v1.1/*",
    "v2": "DELETE"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/message-board",
    "v2": "POST"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/message-board",
    "v2": "GET"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/message-board/*",
    "v2": "PATCH"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/orderfood",
    "v2": "POST"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/orderfood/List",
    "v2": "POST"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/orderfood/excel/personal",
    "v2": "POST"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/account/*",
    "v2": "GET"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/account/change-password/*",
    "v2": "PATCH"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/account/order/*",
    "v2": "GET"
  },
  {
    
    "ptype": "p",
    "v0": "WkCTJ7F4xI",
    "v1": "/hamadan/v1.1/*",
    "v2": "POST"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/menu/List",
    "v2": "POST"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/order/List",
    "v2": "POST"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/order/order-find-store",
    "v2": "POST"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/find-order-list",
    "v2": "POST"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/order",
    "v2": "GET"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/order/*",
    "v2": "GET"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/orderfood/*",
    "v2": "DELETE"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/orderfood/*",
    "v2": "PATCH"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/message-board/List",
    "v2": "POST"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/account/ChangePassword/*",
    "v2": "PATCH"
  },
  {
    
    "ptype": "p",
    "v0": "DbtYHGhUW7",
    "v1": "/hamadan/v1.1/register",
    "v2": "POST"
  },
  {
	"ptype":"p",
	"v0":"DbtYHGhUW7",
	"v1":"/hamadan/v1.1/message-board/*",
	"v2":"DELETE"
  }
]
`
