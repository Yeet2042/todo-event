Repo
// func (r *MongoRepository) Append(ctx context.Context, aggregateID bson.ObjectID, eventType string, payload any) mo.Result[struct{}] {
// 	col, err := r.collection()
// 	if err != nil {
// 		return mo.Err[struct{}](err)
// 	}
// 	raw, err := bson.Marshal(payload)
// 	if err != nil {
// 		return mo.Err[struct{}](err)
// 	}
// 	e := domain.StoredEvent{
// 		ID:          bson.NewObjectID(),
// 		AggregateID: aggregateID,
// 		Type:        eventType,
// 		Payload:     raw,
// 		CreatedAt:   time.Now(),
// 	}
// 	if _, err := col.InsertOne(ctx, e); err != nil {
// 		return mo.Err[struct{}](err)
// 	}
// 	return mo.Ok(struct{}{})
// }

// func (r *MongoRepository) FindByID(ctx context.Context, id bson.ObjectID) mo.Result[domain.Task] {
// 	col, err := r.collection()
// 	if err != nil {
// 		return mo.Err[domain.Task](err)
// 	}
// 	cursor, err := col.Find(ctx,
// 		bson.D{{Key: "aggregate_id", Value: id}},
// 		options.Find().SetSort(bson.D{{Key: "_id", Value: 1}}),
// 	)
// 	if err != nil {
// 		return mo.Err[domain.Task](err)
// 	}
// 	defer cursor.Close(ctx)
// 	var events []domain.StoredEvent
// 	if err := cursor.All(ctx, &events); err != nil {
// 		return mo.Err[domain.Task](err)
// 	}
// 	task, err := domain.Apply(id, events)
// 	if err != nil {
// 		return mo.Err[domain.Task](err)
// 	}
// 	return mo.Ok(task)
// }

// func (r *MongoRepository) FindAll(ctx context.Context) mo.Result[[]domain.Task] {
// 	col, err := r.collection()
// 	if err != nil {
// 		return mo.Err[[]domain.Task](err)
// 	}
// 	pipeline := mongo.Pipeline{
// 		bson.D{{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}},
// 		bson.D{{Key: "$group", Value: bson.D{
// 			{Key: "_id", Value: "$aggregate_id"},
// 			{Key: "events", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
// 		}}},
// 	}
// 	cursor, err := col.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return mo.Err[[]domain.Task](err)
// 	}
// 	defer cursor.Close(ctx)

// 	var tasks []domain.Task
// 	for cursor.Next(ctx) {
// 		var group struct {
// 			AggregateID bson.ObjectID       `bson:"_id"`
// 			Events      []domain.StoredEvent `bson:"events"`
// 		}
// 		if err := cursor.Decode(&group); err != nil {
// 			return mo.Err[[]domain.Task](err)
// 		}
// 		task, err := domain.Apply(group.AggregateID, group.Events)
// 		if err != nil {
// 			return mo.Err[[]domain.Task](err)
// 		}
// 		tasks = append(tasks, task)
// 	}
// 	return mo.Ok(tasks)
// }