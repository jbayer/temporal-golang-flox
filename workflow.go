package greeting

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func SayHelloWorkflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err := workflow.ExecuteActivity(ctx, Greet, name).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	// Invoke another workflow as a child workflow. The parent blocks on
	// .Get until the child completes and returns its result.
	cwo := workflow.ChildWorkflowOptions{
		WorkflowID: "goodbye-child-workflow",
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	var farewell string
	err = workflow.ExecuteChildWorkflow(ctx, SayGoodbyeWorkflow, name).Get(ctx, &farewell)
	if err != nil {
		return "", err
	}

	return result + ". " + farewell, nil
}

// SayGoodbyeWorkflow is invoked by SayHelloWorkflow as a child workflow.
func SayGoodbyeWorkflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err := workflow.ExecuteActivity(ctx, Farewell, name).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}
