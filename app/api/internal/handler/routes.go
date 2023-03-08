// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	martifact "github.com/maolinc/gencode/app/api/internal/handler/martifact"
	martifacttype "github.com/maolinc/gencode/app/api/internal/handler/martifacttype"
	mbug "github.com/maolinc/gencode/app/api/internal/handler/mbug"
	menv "github.com/maolinc/gencode/app/api/internal/handler/menv"
	mexternaldiffrecords "github.com/maolinc/gencode/app/api/internal/handler/mexternaldiffrecords"
	mexternalserver "github.com/maolinc/gencode/app/api/internal/handler/mexternalserver"
	mexternaltask "github.com/maolinc/gencode/app/api/internal/handler/mexternaltask"
	mmasterversion "github.com/maolinc/gencode/app/api/internal/handler/mmasterversion"
	mproject "github.com/maolinc/gencode/app/api/internal/handler/mproject"
	mtypemap "github.com/maolinc/gencode/app/api/internal/handler/mtypemap"
	mversion "github.com/maolinc/gencode/app/api/internal/handler/mversion"
	"github.com/maolinc/gencode/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: martifact.CreateMArtifactHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: martifact.DeleteMArtifactHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: martifact.MArtifactDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/page",
				Handler: martifact.MArtifactPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: martifact.UpdateMArtifactHandler(serverCtx),
			},
		},
		rest.WithPrefix("/artifact/MArtifact"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: martifacttype.CreateMArtifactTypeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: martifacttype.DeleteMArtifactTypeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: martifacttype.MArtifactTypeDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/page",
				Handler: martifacttype.MArtifactTypePageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: martifacttype.UpdateMArtifactTypeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/artifact/MArtifactType"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: mbug.CreateMBugHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: mbug.DeleteMBugHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: mbug.MBugDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/page",
				Handler: mbug.MBugPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: mbug.UpdateMBugHandler(serverCtx),
			},
		},
		rest.WithPrefix("/artifact/MBug"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: menv.CreateMEnvHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: menv.DeleteMEnvHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: menv.MEnvDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/page",
				Handler: menv.MEnvPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: menv.UpdateMEnvHandler(serverCtx),
			},
		},
		rest.WithPrefix("/artifact/MEnv"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: mexternaldiffrecords.CreateMExternalDiffRecordsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: mexternaldiffrecords.DeleteMExternalDiffRecordsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: mexternaldiffrecords.MExternalDiffRecordsDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/page",
				Handler: mexternaldiffrecords.MExternalDiffRecordsPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: mexternaldiffrecords.UpdateMExternalDiffRecordsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/artifact/MExternalDiffRecords"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: mexternalserver.CreateMExternalServerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: mexternalserver.DeleteMExternalServerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: mexternalserver.MExternalServerDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/page",
				Handler: mexternalserver.MExternalServerPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: mexternalserver.UpdateMExternalServerHandler(serverCtx),
			},
		},
		rest.WithPrefix("/artifact/MExternalServer"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: mexternaltask.CreateMExternalTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: mexternaltask.DeleteMExternalTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: mexternaltask.MExternalTaskDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/page",
				Handler: mexternaltask.MExternalTaskPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: mexternaltask.UpdateMExternalTaskHandler(serverCtx),
			},
		},
		rest.WithPrefix("/artifact/MExternalTask"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: mmasterversion.CreateMMasterVersionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: mmasterversion.DeleteMMasterVersionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: mmasterversion.MMasterVersionDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/page",
				Handler: mmasterversion.MMasterVersionPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: mmasterversion.UpdateMMasterVersionHandler(serverCtx),
			},
		},
		rest.WithPrefix("/artifact/MMasterVersion"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: mproject.CreateMProjectHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: mproject.DeleteMProjectHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: mproject.MProjectDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/page",
				Handler: mproject.MProjectPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: mproject.UpdateMProjectHandler(serverCtx),
			},
		},
		rest.WithPrefix("/artifact/MProject"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: mtypemap.CreateMTypeMapHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: mtypemap.DeleteMTypeMapHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: mtypemap.MTypeMapDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/page",
				Handler: mtypemap.MTypeMapPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: mtypemap.UpdateMTypeMapHandler(serverCtx),
			},
		},
		rest.WithPrefix("/artifact/MTypeMap"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: mversion.CreateMVersionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: mversion.DeleteMVersionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: mversion.MVersionDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/page",
				Handler: mversion.MVersionPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: mversion.UpdateMVersionHandler(serverCtx),
			},
		},
		rest.WithPrefix("/artifact/MVersion"),
	)
}
