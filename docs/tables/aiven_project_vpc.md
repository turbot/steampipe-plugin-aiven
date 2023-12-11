---
title: "Steampipe Table: aiven_project_vpc - Query Aiven Project VPCs using SQL"
description: "Allows users to query Aiven Project Virtual Private Clouds (VPCs), providing insights into project-specific VPCs and their configurations."
---

# Table: aiven_project_vpc - Query Aiven Project VPCs using SQL

Aiven Project VPC is a resource within the Aiven platform that allows users to create and manage their own Virtual Private Clouds (VPCs) for their projects. These VPCs provide isolated network environments where users can run their Aiven services, ensuring network security and resource isolation. With Aiven Project VPCs, users have full control over their network environments, including IP range configuration, cloud region selection, and peering connections with other VPCs.

## Table Usage Guide

The `aiven_project_vpc` table provides insights into the Virtual Private Clouds (VPCs) within Aiven projects. As a DevOps engineer, you can explore VPC-specific details through this table, including cloud region, IP range, and peering connections. Utilize it to manage and optimize your network environments, ensuring security and resource isolation for your Aiven services.

## Examples

### Basic info
Explore which projects are active or inactive in your cloud network, and gain insights into their respective network configurations. This information can be useful to manage and optimize your network resources effectively.

```sql+postgres
select
  project_vpc_id,
  project_name,
  state,
  cloud_name,
  network_cidr
from
  aiven_project_vpc;
```

```sql+sqlite
select
  project_vpc_id,
  project_name,
  state,
  cloud_name,
  network_cidr
from
  aiven_project_vpc;
```

### List VPCs with public CIDR blocks
Discover the segments that include VPCs with public CIDR blocks. This can be useful in identifying potential security risks, as these VPCs are not within the private IP address ranges typically used for internal network traffic.

```sql+postgres
select
  project_vpc_id,
  project_name,
  state,
  cloud_name,
  network_cidr
from
  aiven_project_vpc
where
  not network_cidr <<= '10.0.0.0/8'
  and not network_cidr <<= '192.168.0.0/16'
  and not network_cidr <<= '172.16.0.0/12';
```

```sql+sqlite
Error: SQLite does not support CIDR operations.
```

### List VPCs which are not active
Discover the segments that are not currently active within your Virtual Private Clouds (VPCs). This can be useful in identifying inactive resources that could be cleaned up to improve efficiency and reduce costs.

```sql+postgres
select
  project_vpc_id,
  project_name,
  state,
  cloud_name,
  network_cidr
from
  aiven_project_vpc
where
  state <> 'ACTIVE';
```

```sql+sqlite
select
  project_vpc_id,
  project_name,
  state,
  cloud_name,
  network_cidr
from
  aiven_project_vpc
where
  state <> 'ACTIVE';
```

### List VPCs which have invalid peering connections
Explore which Virtual Private Clouds (VPCs) have invalid peering connections to identify potential network issues and ensure efficient data transfer across your projects. This is particularly useful for maintaining secure and reliable network communication within your organization.

```sql+postgres
select
  project_vpc_id,
  project_name,
  state,
  network_cidr
from
  aiven_project_vpc,
  jsonb_array_elements(peering_connections) as c
where
  c ->> 'state' = 'INVALID_SPECIFICATION';
```

```sql+sqlite
select
  project_vpc_id,
  project_name,
  state,
  network_cidr
from
  aiven_project_vpc,
  json_each(peering_connections) as c
where
  json_extract(c.value, '$.state') = 'INVALID_SPECIFICATION';
```

### Get VPC peering connection information of each project
Discover the status and details of Virtual Private Cloud (VPC) peering connections across different projects. This information can be used to understand network configurations and ensure secure and efficient data transfer between projects.

```sql+postgres
select
  project_vpc_id,
  project_name,
  state,
  network_cidr,
  jsonb_pretty(peering_connections) as peering_connections
from
  aiven_project_vpc;
```

```sql+sqlite
select
  project_vpc_id,
  project_name,
  state,
  network_cidr,
  peering_connections
from
  aiven_project_vpc;
```