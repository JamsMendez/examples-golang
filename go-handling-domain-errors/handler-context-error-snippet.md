```
package api

type DomainService interface {
    DoWork(ctx context.Context, userID string) error
}

type API struct {
    domainService DomainService
}

func (a *API) HandleSomeRequest(ctx context.Context, userID string) *Response {
    err := a.domainService.DoWork(ctx, userID)
    if err != nil {
        slog.With(metaerr.GetMetadata(err)...)
            .With("user_id", userID, "error", err)
            .Error("error doing work")

        return NewErrorResponse(err)
    }

    return nil
}



package domain

type Repository interface {
    LocateEntity(ctx context.Context, userID string) (*Entity, error)
    SaveEntity(ctx context.Context, entity *Entity) error
}

type Service struct {
    repo Repository
}

func (s *Service) DoWork(ctx context.Context, userID string) error {
    entity, err := s.repo.LocateEntity(ctx, userID)
    if err != nil {
        return metaerr.WithMetadata(err, "operation", "locate_entity")
    }

    if err := entity.DoWork(); err != nil {
        return metaerr.WithMetadata(err, "operation", "do_work",
            "entity_id", entity.ID)
    }

    if err := s.repo.SaveEntity(ctx, entity); err != nil {
        return metaerr.WithMetadata(err, "operation", "save_entity",
            "entity_id", entity.ID)
    }

    return nil
}



```
