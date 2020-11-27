// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build linux

package probe

import (
	"github.com/DataDog/datadog-agent/pkg/security/ebpf"
)

type activeApprover = activeKFilter
type activeApprovers = activeKFilters

func approveBasename(tableName string, eventType EventType, basename string) (activeApprover, error) {
	return &mapEventMask{
		tableName: tableName,
		key:       basename,
		tableKey:  ebpf.NewStringMapItem(basename, BasenameFilterSize),
		eventMask: uint64(1 << (eventType - 1)),
	}, nil
}

func approveBasenames(tableName string, eventType EventType, basenames ...string) (approvers []activeApprover, _ error) {
	for _, basename := range basenames {
		activeApprover, err := approveBasename(tableName, eventType, basename)
		if err != nil {
			return nil, err
		}
		approvers = append(approvers, activeApprover)
	}
	return approvers, nil
}

func setFlagsFilter(tableName string, flags ...int) (activeApprover, error) {
	var flagsItem ebpf.Uint32MapItem

	for _, flag := range flags {
		flagsItem |= ebpf.Uint32MapItem(flag)
	}

	if flagsItem != 0 {
		return &arrayEntry{
			tableName: tableName,
			index:     uint32(0),
			value:     flagsItem,
			zeroValue: ebpf.ZeroUint32MapItem,
		}, nil
	}

	return nil, nil
}

func approveFlags(tableName string, flags ...int) (activeApprover, error) {
	return setFlagsFilter(tableName, flags...)
}
