col username format a15
col machine format a50
col owner format a12
col type format a25
col object format a25

set pages 50
set lines 150


---------------------------------------------------
-- Specfic sqls
---------------------------------------------------
-- ora_exadata_cluster -> ora_phy_db [taillabel="hosts"]
-- TBD --> get from OEM

-- ora_phy_db -> ora_schema [taillabel="hosts"]
select distinct a.value ora_phy_db, b.owner ora_schema from v$parameter a, dba_tables b where a.name = 'db_unique_name';

-- ora_table -> ora_schema [taillabel="member_of"]
select a.value ora_phy_db, b.owner ora_schema, b.table_name from v$parameter a, dba_tables b where a.name = 'db_unique_name';

-- ora_session -> ora_lsnr [taillabel="connected_to"]
-- TBD --> may not be possible except from listener log files on db hosts

-- ora_session -> ora_user [taillabel="authenticated_as"]
-- old style: select a.value ora_phy_db, b.sid || ':' || b.serial# ora_session, b.username ora_user from v$parameter a, gv$session b where a.name = 'db_unique_name';
select a.value ora_phy_db, b.con_id || ':' || inst_id || ':' ||  b.sid || ':' || b.serial# ora_session, b.username ora_user from v$parameter a, gv$session b where a.name = 'db_unique_name';

-- ora_session -> ora_table [taillabel="interacts_with"]
select a.value ora_phy_db, b.con_id || ':' || inst_id || ':' || b.sid || ':' || b.serial# ora_session from v$parameter a, gv$session b, gv$access c where a.name = 'db_unique_name' and b.con_id = c.con_id and b.inst_id = c.inst_id and b.sid = c.sid;

-- tap_instance -> ora_session [taillabel="connected_to"]
select a.value ora_phy_db, b.con_id || ':' || inst_id || ':' ||  b.sid || ':' || b.serial# ora_session, b.machine tap_instance from v$parameter a, gv$session b where a.name = 'db_unique_name';

-- Reference: Unique with con_id, inst_id, sid, serial#
--select con_id || ':' || inst_id || ':' || sid || ':' || serial# from gv$session where rownum < 20;

---------------------------------------------------
-- General samples - with json output from Oracle
---------------------------------------------------

-- DB Name
select json_object('DbName' is name) database_name from v$database;

-- Database 'Unique' name
select json_object('DbUniqueName' is a.value) from v$parameter a where name = 'db_unique_name';

-- The db-in-db 'pluggable' database name
select json_object('PdbName' is global_name) from global_name;

-- TAP or other machine connecting to the database
select json_object('Machine' is b.machine) from gv$access a, gv$session b where a.sid=b.sid and rownum < 10 order by b.machine;

-- TAP or other machine, and which object it uses (like table, index - here, filtered to give only tables)
-- SocketPort is not the same as Listener port - I guess it is the port of the app machine
select json_object('AppMachine' is b.machine, 'ObjectOwner' is a.owner, 'ObjectType' is a.type, 'Object' is a.object, 'AppUser' is b.username, 'SocketPort' is b.port, 'Server' is server) from gv$access a, gv$session b where a.sid=b.sid and a.type like 'TABLE%' and a.owner like '%MGR'  order by b.machine, b.username, a.type, a.owner, a.object ;

-- Just the listener
select json_object('Listener' is a.value) from gv$parameter a where name like 'local_listener';

-- Hosts of Oracle Cluster
select json_object('ClusterHost' is a.host_name) from gv$instance a;

-- Instance of multi-instance RAC cluster and its host
select json_object('DbInstance' is a.instance_name, 'ClusterHost' is a.host_name) from gv$instance a;
