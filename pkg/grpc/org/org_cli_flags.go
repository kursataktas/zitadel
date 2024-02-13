// Code generated by protoc-gen-cli-client. DO NOT EDIT.

package org

import (
	cli_client "github.com/adlerhurst/cli-client"
	pflag "github.com/spf13/pflag"
	object "github.com/zitadel/zitadel/pkg/grpc/object"
	os "os"
)

type DomainFlag struct {
	*Domain

	changed bool
	set     *pflag.FlagSet

	orgIdFlag          *cli_client.StringParser
	detailsFlag        *object.ObjectDetailsFlag
	domainNameFlag     *cli_client.StringParser
	isVerifiedFlag     *cli_client.BoolParser
	isPrimaryFlag      *cli_client.BoolParser
	validationTypeFlag *cli_client.EnumParser[DomainValidationType]
}

func (x *DomainFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("Domain", pflag.ContinueOnError)

	x.orgIdFlag = cli_client.NewStringParser(x.set, "org-id", "")
	x.domainNameFlag = cli_client.NewStringParser(x.set, "domain-name", "")
	x.isVerifiedFlag = cli_client.NewBoolParser(x.set, "is-verified", "")
	x.isPrimaryFlag = cli_client.NewBoolParser(x.set, "is-primary", "")
	x.validationTypeFlag = cli_client.NewEnumParser[DomainValidationType](x.set, "validation-type", "")
	x.detailsFlag = &object.ObjectDetailsFlag{ObjectDetails: new(object.ObjectDetails)}
	x.detailsFlag.AddFlags(x.set)
	parent.AddFlagSet(x.set)
}

func (x *DomainFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := cli_client.FieldIndexes(args, "details")

	if err := x.set.Parse(flagIndexes.Primitives().Args); err != nil {
		cli_client.Logger().Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if flagIdx := flagIndexes.LastByName("details"); flagIdx != nil {
		x.detailsFlag.ParseFlags(x.set, flagIdx.Args)
	}

	if x.orgIdFlag.Changed() {
		x.changed = true
		x.Domain.OrgId = *x.orgIdFlag.Value
	}

	if x.detailsFlag.Changed() {
		x.changed = true
		x.Domain.Details = x.detailsFlag.ObjectDetails
	}

	if x.domainNameFlag.Changed() {
		x.changed = true
		x.Domain.DomainName = *x.domainNameFlag.Value
	}
	if x.isVerifiedFlag.Changed() {
		x.changed = true
		x.Domain.IsVerified = *x.isVerifiedFlag.Value
	}
	if x.isPrimaryFlag.Changed() {
		x.changed = true
		x.Domain.IsPrimary = *x.isPrimaryFlag.Value
	}
	if x.validationTypeFlag.Changed() {
		x.changed = true
		x.Domain.ValidationType = *x.validationTypeFlag.Value
	}
}

func (x *DomainFlag) Changed() bool {
	return x.changed
}

type DomainNameQueryFlag struct {
	*DomainNameQuery

	changed bool
	set     *pflag.FlagSet

	nameFlag   *cli_client.StringParser
	methodFlag *cli_client.EnumParser[object.TextQueryMethod]
}

func (x *DomainNameQueryFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("DomainNameQuery", pflag.ContinueOnError)

	x.nameFlag = cli_client.NewStringParser(x.set, "name", "")
	x.methodFlag = cli_client.NewEnumParser[object.TextQueryMethod](x.set, "method", "")
	parent.AddFlagSet(x.set)
}

func (x *DomainNameQueryFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := cli_client.FieldIndexes(args)

	if err := x.set.Parse(flagIndexes.Primitives().Args); err != nil {
		cli_client.Logger().Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if x.nameFlag.Changed() {
		x.changed = true
		x.DomainNameQuery.Name = *x.nameFlag.Value
	}
	if x.methodFlag.Changed() {
		x.changed = true
		x.DomainNameQuery.Method = *x.methodFlag.Value
	}
}

func (x *DomainNameQueryFlag) Changed() bool {
	return x.changed
}

type DomainSearchQueryFlag struct {
	*DomainSearchQuery

	changed bool
	set     *pflag.FlagSet

	domainNameQueryFlag *DomainNameQueryFlag
}

func (x *DomainSearchQueryFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("DomainSearchQuery", pflag.ContinueOnError)

	x.domainNameQueryFlag = &DomainNameQueryFlag{DomainNameQuery: new(DomainNameQuery)}
	x.domainNameQueryFlag.AddFlags(x.set)
	parent.AddFlagSet(x.set)
}

func (x *DomainSearchQueryFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := cli_client.FieldIndexes(args, "domain-name-query")

	if err := x.set.Parse(flagIndexes.Primitives().Args); err != nil {
		cli_client.Logger().Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if flagIdx := flagIndexes.LastByName("domain-name-query"); flagIdx != nil {
		x.domainNameQueryFlag.ParseFlags(x.set, flagIdx.Args)
	}

	switch cli_client.FieldIndexes(args, "domain-name-query").Last().Flag {
	case "domain-name-query":
		if x.domainNameQueryFlag.Changed() {
			x.changed = true
			x.DomainSearchQuery.Query = &DomainSearchQuery_DomainNameQuery{DomainNameQuery: x.domainNameQueryFlag.DomainNameQuery}
		}
	}
}

func (x *DomainSearchQueryFlag) Changed() bool {
	return x.changed
}

type OrgFlag struct {
	*Org

	changed bool
	set     *pflag.FlagSet

	idFlag            *cli_client.StringParser
	detailsFlag       *object.ObjectDetailsFlag
	stateFlag         *cli_client.EnumParser[OrgState]
	nameFlag          *cli_client.StringParser
	primaryDomainFlag *cli_client.StringParser
}

func (x *OrgFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("Org", pflag.ContinueOnError)

	x.idFlag = cli_client.NewStringParser(x.set, "id", "")
	x.stateFlag = cli_client.NewEnumParser[OrgState](x.set, "state", "")
	x.nameFlag = cli_client.NewStringParser(x.set, "name", "")
	x.primaryDomainFlag = cli_client.NewStringParser(x.set, "primary-domain", "")
	x.detailsFlag = &object.ObjectDetailsFlag{ObjectDetails: new(object.ObjectDetails)}
	x.detailsFlag.AddFlags(x.set)
	parent.AddFlagSet(x.set)
}

func (x *OrgFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := cli_client.FieldIndexes(args, "details")

	if err := x.set.Parse(flagIndexes.Primitives().Args); err != nil {
		cli_client.Logger().Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if flagIdx := flagIndexes.LastByName("details"); flagIdx != nil {
		x.detailsFlag.ParseFlags(x.set, flagIdx.Args)
	}

	if x.idFlag.Changed() {
		x.changed = true
		x.Org.Id = *x.idFlag.Value
	}

	if x.detailsFlag.Changed() {
		x.changed = true
		x.Org.Details = x.detailsFlag.ObjectDetails
	}

	if x.stateFlag.Changed() {
		x.changed = true
		x.Org.State = *x.stateFlag.Value
	}
	if x.nameFlag.Changed() {
		x.changed = true
		x.Org.Name = *x.nameFlag.Value
	}
	if x.primaryDomainFlag.Changed() {
		x.changed = true
		x.Org.PrimaryDomain = *x.primaryDomainFlag.Value
	}
}

func (x *OrgFlag) Changed() bool {
	return x.changed
}

type OrgDomainQueryFlag struct {
	*OrgDomainQuery

	changed bool
	set     *pflag.FlagSet

	domainFlag *cli_client.StringParser
	methodFlag *cli_client.EnumParser[object.TextQueryMethod]
}

func (x *OrgDomainQueryFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("OrgDomainQuery", pflag.ContinueOnError)

	x.domainFlag = cli_client.NewStringParser(x.set, "domain", "")
	x.methodFlag = cli_client.NewEnumParser[object.TextQueryMethod](x.set, "method", "")
	parent.AddFlagSet(x.set)
}

func (x *OrgDomainQueryFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := cli_client.FieldIndexes(args)

	if err := x.set.Parse(flagIndexes.Primitives().Args); err != nil {
		cli_client.Logger().Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if x.domainFlag.Changed() {
		x.changed = true
		x.OrgDomainQuery.Domain = *x.domainFlag.Value
	}
	if x.methodFlag.Changed() {
		x.changed = true
		x.OrgDomainQuery.Method = *x.methodFlag.Value
	}
}

func (x *OrgDomainQueryFlag) Changed() bool {
	return x.changed
}

type OrgNameQueryFlag struct {
	*OrgNameQuery

	changed bool
	set     *pflag.FlagSet

	nameFlag   *cli_client.StringParser
	methodFlag *cli_client.EnumParser[object.TextQueryMethod]
}

func (x *OrgNameQueryFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("OrgNameQuery", pflag.ContinueOnError)

	x.nameFlag = cli_client.NewStringParser(x.set, "name", "")
	x.methodFlag = cli_client.NewEnumParser[object.TextQueryMethod](x.set, "method", "")
	parent.AddFlagSet(x.set)
}

func (x *OrgNameQueryFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := cli_client.FieldIndexes(args)

	if err := x.set.Parse(flagIndexes.Primitives().Args); err != nil {
		cli_client.Logger().Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if x.nameFlag.Changed() {
		x.changed = true
		x.OrgNameQuery.Name = *x.nameFlag.Value
	}
	if x.methodFlag.Changed() {
		x.changed = true
		x.OrgNameQuery.Method = *x.methodFlag.Value
	}
}

func (x *OrgNameQueryFlag) Changed() bool {
	return x.changed
}

type OrgQueryFlag struct {
	*OrgQuery

	changed bool
	set     *pflag.FlagSet

	nameQueryFlag   *OrgNameQueryFlag
	domainQueryFlag *OrgDomainQueryFlag
	stateQueryFlag  *OrgStateQueryFlag
}

func (x *OrgQueryFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("OrgQuery", pflag.ContinueOnError)

	x.nameQueryFlag = &OrgNameQueryFlag{OrgNameQuery: new(OrgNameQuery)}
	x.nameQueryFlag.AddFlags(x.set)
	x.domainQueryFlag = &OrgDomainQueryFlag{OrgDomainQuery: new(OrgDomainQuery)}
	x.domainQueryFlag.AddFlags(x.set)
	x.stateQueryFlag = &OrgStateQueryFlag{OrgStateQuery: new(OrgStateQuery)}
	x.stateQueryFlag.AddFlags(x.set)
	parent.AddFlagSet(x.set)
}

func (x *OrgQueryFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := cli_client.FieldIndexes(args, "name-query", "domain-query", "state-query")

	if err := x.set.Parse(flagIndexes.Primitives().Args); err != nil {
		cli_client.Logger().Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if flagIdx := flagIndexes.LastByName("name-query"); flagIdx != nil {
		x.nameQueryFlag.ParseFlags(x.set, flagIdx.Args)
	}

	if flagIdx := flagIndexes.LastByName("domain-query"); flagIdx != nil {
		x.domainQueryFlag.ParseFlags(x.set, flagIdx.Args)
	}

	if flagIdx := flagIndexes.LastByName("state-query"); flagIdx != nil {
		x.stateQueryFlag.ParseFlags(x.set, flagIdx.Args)
	}

	switch cli_client.FieldIndexes(args, "name-query", "domain-query", "state-query").Last().Flag {
	case "name-query":
		if x.nameQueryFlag.Changed() {
			x.changed = true
			x.OrgQuery.Query = &OrgQuery_NameQuery{NameQuery: x.nameQueryFlag.OrgNameQuery}
		}
	case "domain-query":
		if x.domainQueryFlag.Changed() {
			x.changed = true
			x.OrgQuery.Query = &OrgQuery_DomainQuery{DomainQuery: x.domainQueryFlag.OrgDomainQuery}
		}
	case "state-query":
		if x.stateQueryFlag.Changed() {
			x.changed = true
			x.OrgQuery.Query = &OrgQuery_StateQuery{StateQuery: x.stateQueryFlag.OrgStateQuery}
		}
	}
}

func (x *OrgQueryFlag) Changed() bool {
	return x.changed
}

type OrgStateQueryFlag struct {
	*OrgStateQuery

	changed bool
	set     *pflag.FlagSet

	stateFlag *cli_client.EnumParser[OrgState]
}

func (x *OrgStateQueryFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("OrgStateQuery", pflag.ContinueOnError)

	x.stateFlag = cli_client.NewEnumParser[OrgState](x.set, "state", "")
	parent.AddFlagSet(x.set)
}

func (x *OrgStateQueryFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := cli_client.FieldIndexes(args)

	if err := x.set.Parse(flagIndexes.Primitives().Args); err != nil {
		cli_client.Logger().Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if x.stateFlag.Changed() {
		x.changed = true
		x.OrgStateQuery.State = *x.stateFlag.Value
	}
}

func (x *OrgStateQueryFlag) Changed() bool {
	return x.changed
}