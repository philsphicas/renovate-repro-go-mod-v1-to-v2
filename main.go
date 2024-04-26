package main

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

func main() {
	fmt.Println(armkubernetesconfiguration.AKSIdentityTypeSystemAssigned)
}
