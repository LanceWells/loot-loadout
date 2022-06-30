package http

// import (
// 	"log"

// 	"github.com/labstack/echo/v4"
// 	api "github.com/lantspants/lootloadout/api/characterimage/v1"
// 	service "github.com/lantspants/lootloadout/backend/characterimage/characterimage"
// )

// // Ensure that this type implements the underlying interface.
// var _ api.ServerInterface = CharacterImageServer{}

// type CharacterImageServer struct {
// 	l *log.Logger
// 	s service.CharacterImageService
// }

// func NewCharacterImageServer(
// 	l *log.Logger,
// 	s service.CharacterImageService,
// ) *CharacterImageServer {
// 	return &CharacterImageServer{
// 		l: l,
// 		s: s,
// 	}
// }

// // (GET /body)
// func (r CharacterImageServer) ListBodies(ctx echo.Context) error {
// 	return nil
// }

// // (POST /body)
// func (r CharacterImageServer) AddBody(ctx echo.Context) error {
// 	return nil
// }

// // (GET /body/{bodyID}/animation)
// func (r CharacterImageServer) ListAnimations(ctx echo.Context, bodyID string) error {
// 	return nil
// }

// // (POST /body/{bodyID}/animation)
// func (r CharacterImageServer) AddAnimation(ctx echo.Context, bodyID string) error {
// 	return nil
// }

// // (POST /body/{bodyID}/animation/{animationID}/frame)
// func (r CharacterImageServer) AddAnimationFrame(ctx echo.Context, bodyID string, animationID string) error {
// 	return nil
// }

// // (GET /body/{bodyID}/dynamic)
// func (r CharacterImageServer) ListDyanmics(ctx echo.Context, bodyID string) error {
// 	return nil
// }

// // (POST /body/{bodyID}/dynamic)
// func (r CharacterImageServer) AddDynamic(ctx echo.Context, bodyID string) error {
// 	return nil
// }

// // (POST /body/{bodyID}/dynamicMapping)
// func (r CharacterImageServer) AddDynamicMapping(ctx echo.Context, bodyID string) error {
// 	return nil
// }

// // (GET /body/{bodyID}/static)
// func (r CharacterImageServer) ListStatics(ctx echo.Context, bodyID string) error {
// 	return nil
// }

// // (POST /body/{bodyID}/static)
// func (r CharacterImageServer) AddStatic(ctx echo.Context, bodyID string) error {
// 	return nil
// }

// // (GET /prop)
// func (r CharacterImageServer) ListProps(ctx echo.Context) error {
// 	return nil
// }

// // (POST /prop)
// func (r CharacterImageServer) AddProp(ctx echo.Context) error {
// 	return nil
// }
