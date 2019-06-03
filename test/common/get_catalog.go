package common

import (
	"context"
	"testing"

	openapi "github.com/openservicebrokerapi/osb-checker/autogenerated/go-client"
	. "github.com/openservicebrokerapi/osb-checker/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCatalog(
	t *testing.T,
) {
	Convey("Query service catalog", t, func() {

		Convey("should return 412 PreconditionFailed if missing X-Broker-API-Version header", func() {
			_, resp, err := cli.CatalogApi.CatalogGet(
				authCtx, "", &openapi.CatalogGetOpts{})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 412)
		})

		if CONF.Authentication.AuthType != TypeNoauth {
			Convey("should return 401 Unauthorized if missing Authorization header", func() {
				_, resp, err := cli.CatalogApi.CatalogGet(
					context.Background(), CONF.APIVersion, &openapi.CatalogGetOpts{})

				So(err, ShouldNotBeNil)
				So(resp.StatusCode, ShouldEqual, 401)
			})
		}

		Convey("should return list of registered service classes as JSON payload", func() {
			_, resp, err := cli.CatalogApi.CatalogGet(
				authCtx, CONF.APIVersion, &openapi.CatalogGetOpts{})

			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, 200)
		})
	})
}
