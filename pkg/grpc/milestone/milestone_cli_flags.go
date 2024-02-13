// Code generated by protoc-gen-cli-client. DO NOT EDIT.

package milestone

import (
	cli_client "github.com/adlerhurst/cli-client"
	pflag "github.com/spf13/pflag"
	os "os"
)

type IsReachedQueryFlag struct {
	*IsReachedQuery

	changed bool
	set     *pflag.FlagSet

	reachedFlag *cli_client.BoolParser
}

func (x *IsReachedQueryFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("IsReachedQuery", pflag.ContinueOnError)

	x.reachedFlag = cli_client.NewBoolParser(x.set, "reached", "")
	parent.AddFlagSet(x.set)
}

func (x *IsReachedQueryFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := cli_client.FieldIndexes(args)

	if err := x.set.Parse(flagIndexes.Primitives().Args); err != nil {
		cli_client.Logger().Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if x.reachedFlag.Changed() {
		x.changed = true
		x.IsReachedQuery.Reached = *x.reachedFlag.Value
	}
}

func (x *IsReachedQueryFlag) Changed() bool {
	return x.changed
}

type MilestoneFlag struct {
	*Milestone

	changed bool
	set     *pflag.FlagSet

	typeFlag        *cli_client.EnumParser[MilestoneType]
	reachedDateFlag *cli_client.TimestampParser
}

func (x *MilestoneFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("Milestone", pflag.ContinueOnError)

	x.typeFlag = cli_client.NewEnumParser[MilestoneType](x.set, "type", "")
	x.reachedDateFlag = cli_client.NewTimestampParser(x.set, "reached-date", "")
	parent.AddFlagSet(x.set)
}

func (x *MilestoneFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := cli_client.FieldIndexes(args)

	if err := x.set.Parse(flagIndexes.Primitives().Args); err != nil {
		cli_client.Logger().Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if x.typeFlag.Changed() {
		x.changed = true
		x.Milestone.Type = *x.typeFlag.Value
	}
	if x.reachedDateFlag.Changed() {
		x.changed = true
		x.Milestone.ReachedDate = x.reachedDateFlag.Value
	}
}

func (x *MilestoneFlag) Changed() bool {
	return x.changed
}

type MilestoneQueryFlag struct {
	*MilestoneQuery

	changed bool
	set     *pflag.FlagSet

	isReachedQueryFlag *IsReachedQueryFlag
}

func (x *MilestoneQueryFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("MilestoneQuery", pflag.ContinueOnError)

	x.isReachedQueryFlag = &IsReachedQueryFlag{IsReachedQuery: new(IsReachedQuery)}
	x.isReachedQueryFlag.AddFlags(x.set)
	parent.AddFlagSet(x.set)
}

func (x *MilestoneQueryFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := cli_client.FieldIndexes(args, "is-reached-query")

	if err := x.set.Parse(flagIndexes.Primitives().Args); err != nil {
		cli_client.Logger().Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if flagIdx := flagIndexes.LastByName("is-reached-query"); flagIdx != nil {
		x.isReachedQueryFlag.ParseFlags(x.set, flagIdx.Args)
	}

	switch cli_client.FieldIndexes(args, "is-reached-query").Last().Flag {
	case "is-reached-query":
		if x.isReachedQueryFlag.Changed() {
			x.changed = true
			x.MilestoneQuery.Query = &MilestoneQuery_IsReachedQuery{IsReachedQuery: x.isReachedQueryFlag.IsReachedQuery}
		}
	}
}

func (x *MilestoneQueryFlag) Changed() bool {
	return x.changed
}