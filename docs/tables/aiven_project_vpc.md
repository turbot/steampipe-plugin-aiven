# Table: aiven_project_vpc

Represents virtual network in Projects.

## Examples

### Basic info

```sql
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

```sql
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

### List VPCs which are not active

```sql
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

### List VPCs which has invalid peering connections

```sql
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

### Get VPC peering connections information for each project

```sql
select
  project_vpc_id,
  project_name,
  state,
  network_cidr,
  jsonb_pretty(peering_connections) as peering_connections
from
  aiven_project_vpc;
```
