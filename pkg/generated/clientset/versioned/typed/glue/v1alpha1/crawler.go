/*
Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License"). You may
not use this file except in compliance with the License. A copy of the
License is located at

     http://aws.amazon.com/apache2.0/

or in the "license" file accompanying this file. This file is distributed
on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing
permissions and limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "awsoperator.io/pkg/apis/glue/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CrawlersGetter has a method to return a CrawlerInterface.
// A group's client should implement this interface.
type CrawlersGetter interface {
	Crawlers(namespace string) CrawlerInterface
}

// CrawlerInterface has methods to work with Crawler resources.
type CrawlerInterface interface {
	Create(*v1alpha1.Crawler) (*v1alpha1.Crawler, error)
	Update(*v1alpha1.Crawler) (*v1alpha1.Crawler, error)
	UpdateStatus(*v1alpha1.Crawler) (*v1alpha1.Crawler, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Crawler, error)
	List(opts v1.ListOptions) (*v1alpha1.CrawlerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Crawler, err error)
	CrawlerExpansion
}

// crawlers implements CrawlerInterface
type crawlers struct {
	client rest.Interface
	ns     string
}

// newCrawlers returns a Crawlers
func newCrawlers(c *GlueV1alpha1Client, namespace string) *crawlers {
	return &crawlers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the crawler, and returns the corresponding crawler object, and an error if there is any.
func (c *crawlers) Get(name string, options v1.GetOptions) (result *v1alpha1.Crawler, err error) {
	result = &v1alpha1.Crawler{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("crawlers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Crawlers that match those selectors.
func (c *crawlers) List(opts v1.ListOptions) (result *v1alpha1.CrawlerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CrawlerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("crawlers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested crawlers.
func (c *crawlers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("crawlers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a crawler and creates it.  Returns the server's representation of the crawler, and an error, if there is any.
func (c *crawlers) Create(crawler *v1alpha1.Crawler) (result *v1alpha1.Crawler, err error) {
	result = &v1alpha1.Crawler{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("crawlers").
		Body(crawler).
		Do().
		Into(result)
	return
}

// Update takes the representation of a crawler and updates it. Returns the server's representation of the crawler, and an error, if there is any.
func (c *crawlers) Update(crawler *v1alpha1.Crawler) (result *v1alpha1.Crawler, err error) {
	result = &v1alpha1.Crawler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("crawlers").
		Name(crawler.Name).
		Body(crawler).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *crawlers) UpdateStatus(crawler *v1alpha1.Crawler) (result *v1alpha1.Crawler, err error) {
	result = &v1alpha1.Crawler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("crawlers").
		Name(crawler.Name).
		SubResource("status").
		Body(crawler).
		Do().
		Into(result)
	return
}

// Delete takes name of the crawler and deletes it. Returns an error if one occurs.
func (c *crawlers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("crawlers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *crawlers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("crawlers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched crawler.
func (c *crawlers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Crawler, err error) {
	result = &v1alpha1.Crawler{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("crawlers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
