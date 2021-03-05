Config
==================================

.. toctree::
   :maxdepth: 2
   :caption: Contents:

========

Configuring ``Whitetail`` takes the form of editing the ``config.json`` file which is contained in the whitetail configmap

Basic Configuration
-------------------

The basic configuration handles the server itself that `Whitetail` runs. This includes the following values:

+-----------+------------------------------------------------+
| Name      | Description                                    |
+===========+================================================+
| http-port | The port to serve the `Whitetail` UI on        |
+-----------+------------------------------------------------+
| tcp-port  | The port to listen to TCP logs on              |
+-----------+------------------------------------------------+
| udp-port  | The port to listen to UDP logs on              |
+-----------+------------------------------------------------+
| basepath  | the basepath to serve the various endpoints on |
+-----------+------------------------------------------------+

Database
--------

The database configuration is held in the ``database`` key in the configuration file. This defines which database ``Whitetail`` will use to hold log and index information

.. code-block:: JSON

    {
        "database": {
            "url": "http://localhost:9090"
        },
    }

By default the URL points to a local instsance of ``Ceres``, which works for the Kubernetes deployment. You can instead point it at an external instsance via the URL if you wish.

Logging
-------

Logging configuration mainly concerns itself with the cleanup process to remove old logs, however it does also configure some aspects of the log message formatting.

+-----------------------+--------------------------------------------------------------------------------------------------------------------------------------+
| Name                  | Description                                                                                                                          |
+=======================+======================================================================================================================================+
| max-age-days          | How many days to keep logs for (integer)                                                                                             |
+-----------------------+--------------------------------------------------------------------------------------------------------------------------------------+
| poll-rate             | How often to check for old logs. Is of the form `< number >< time unit >` where valid time units are `ns`, `us`, `ms`, `s`, `m`, `h` |
+-----------------------+--------------------------------------------------------------------------------------------------------------------------------------+
| concise-logger        | Should the logger name be compaceted for ease of viewing (bool)                                                                      |
+-----------------------+--------------------------------------------------------------------------------------------------------------------------------------+
| hoverable-long-logger | Should the logger name be expanded when you hover over it (bool)                                                                     |
+-----------------------+--------------------------------------------------------------------------------------------------------------------------------------+

Branding
--------

Branding configuration allows for ``Whitetail`` to be customized to fit your product that it is being used in conjunction with. You an either change these through the ``Settings`` page in the UI or through the configuration file.

+----------------------------+-------------------------------------------------+
| Name                       | Description                                     |
+----------------------------+-------------------------------------------------+
| primary_color.background`  | Primary branding color                          |
+============================+=================================================+
| primary_color.text         | Color for text over primary branding color      |
+----------------------------+-------------------------------------------------+
| secondary_color.background | Secondary branding color                        |
+----------------------------+-------------------------------------------------+
| secondary_color.text       | Color for text over secondary branding color    |
+----------------------------+-------------------------------------------------+
| tertiary_color.backgroud   | Tertiary branding color                         |
+----------------------------+-------------------------------------------------+
| tertiary_color.text        | Color for text over tertiary branding color     |
+----------------------------+-------------------------------------------------+
| INFO_color                 | Color to be used to highligh `INFO` level logs  |
+----------------------------+-------------------------------------------------+
| WARN_color                 | Color to be used to highligh `WARN` level logs  |
+----------------------------+-------------------------------------------------+
| DEBUG_color                | Color to be used to highligh `DEBUG` level logs |
+----------------------------+-------------------------------------------------+
| TRACE_color                | Color to be used to highligh `TRACE` level logs |
+----------------------------+-------------------------------------------------+
| ERROR_color                | Color to be used to highligh `ERROR` level logs |
+----------------------------+-------------------------------------------------+

In addition, you can configure the logo shown in the UI by placing your own logo file at ``< whitetail root >/config/custom/logo/logo.png`` and you can configure the icon shown in the browser by placing your own icon file at ``< whitetail root >/config/custom/icon/favicon.png``