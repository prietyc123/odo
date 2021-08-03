package project

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/odo/tests/helper"
	"github.com/tidwall/gjson"
)

var _ = Describe("[rfe_id:100][crit:medium][vendor:cnv-qe@redhat.com][level:component]odo project command tests", func() {
	var commonVar helper.CommonVar

	// This is run before every Spec (It)
	var _ = BeforeEach(func() {
		commonVar = helper.CommonBeforeEach()
	})

	// Clean up after the test
	// This is run after every Spec (It)
	var _ = AfterEach(func() {
		helper.CommonAfterEach(commonVar)
	})

	It("should display machine ouptut when getting help for odo project list [test_id:1000]", func() {
		output := helper.Cmd("odo", "project", "list", "--help").ShouldPass().Out()
		Expect(output).To(ContainSubstring("Specify output format, supported format: json"))
	})

	It("should be able to get project [test_id:1001]", func() {
		projectGetJSON := helper.Cmd("odo", "project", "get", "-o", "json").ShouldPass().Out()
		getOutputJSON, err := helper.Unindented(projectGetJSON)
		Expect(err).Should(BeNil())
		valuesJSON := gjson.GetMany(getOutputJSON, "kind", "status.active")
		expectedJSON := []string{"Project", "true"}
		Expect(helper.GjsonMatcher(valuesJSON, expectedJSON)).To(Equal(true))

	})

	It("should display the help of project command [test_id:1002]", func() {
		projectHelp := helper.Cmd("odo", "project", "-h").ShouldPass().Out()
		Expect(projectHelp).To(ContainSubstring("Perform project operations"))
	})

	It("should display only the project name when running command with -q flag [test_id:1003]", func() {
		projectName := helper.Cmd("odo", "project", "get", "-q").ShouldPass().Out()
		Expect(projectName).Should(Equal(commonVar.Project))
	})

	It("should list current empty project in json format [test_id:1004]", func() {
		projectListJSON := helper.Cmd("odo", "project", "list", "-o", "json").ShouldPass().Out()
		listOutputJSON, err := helper.Unindented(projectListJSON)
		Expect(err).Should(BeNil())
		partOfProjectListJSON, err := helper.Unindented(`{"kind":"Project","apiVersion":"odo.dev/v1alpha1","metadata":{"name":"` + commonVar.Project + `","namespace":"` + commonVar.Project + `","creationTimestamp":null},"spec":{},"status":{"active":true}}`)
		Expect(err).Should(BeNil())
		Expect(listOutputJSON).To(ContainSubstring(partOfProjectListJSON))
	})

	It("should list current empty project [test_id:1005]", func() {
		helper.WaitForCmdOut("odo", []string{"project", "list"}, 1, true, func(output string) bool {
			return strings.Contains(output, commonVar.Project)
		})
	})

	When("creating a new project", func() {
		var projectName string

		BeforeEach(func() {
			projectName = helper.RandString(6)
			helper.Cmd("odo", "project", "create", projectName).ShouldPass()
		})

		It("should delete a project with --wait [test_id:1006]", func() {
			output := helper.Cmd("odo", "project", "delete", projectName, "-f", "--wait").ShouldPass().Out()
			Expect(output).To(ContainSubstring("Waiting for project to be deleted"))
		})
	})

	When("creating a new project with -o json", func() {
		var projectName string

		BeforeEach(func() {
			projectName = helper.RandString(6)
			helper.Cmd("odo", "project", "create", projectName, "-o", "json").ShouldPass()
		})

		It("should delete project and show output in json format [test_id:1007]", func() {
			actual := helper.Cmd("odo", "project", "delete", projectName, "-o", "json").ShouldPass().Out()
			values := gjson.GetMany(actual, "kind", "message")
			expected := []string{"Project", "Deleted project :"}
			Expect(helper.GjsonMatcher(values, expected)).To(Equal(true))

		})
	})
})
