/*
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package io.fabric8.kubernetes.api.model;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class ListOptionsTest {
  @Test
  void testBuilder() {
    ListOptions listOptions = new io.fabric8.kubernetes.api.model.ListOptionsBuilder()
        .withLimit(100L)
        .withContinue("23243434")
        .withFieldSelector("metadata.name=my-service")
        .build();

    assertEquals(100L, listOptions.getLimit());
    assertEquals("23243434", listOptions.getContinue());
    assertEquals("metadata.name=my-service", listOptions.getFieldSelector());
  }
}
