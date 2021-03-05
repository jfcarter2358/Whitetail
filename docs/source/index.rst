Welcome to Whitetail's documentation!
==================================

.. toctree::
   :maxdepth: 2
   :caption: Contents:

   About.rst
   Config.rst
   AQL.rst

.. _Address:

Indices and tables
==================

* :ref:`genindex`
* :ref:`modindex`
* :ref:`search`

========

Installation
------------

To run Whitetail inside Kubernetes, apply any of the following combinations of manifests inside the ``manifests`` directory

No PVCs

- configmap.ceres.yaml
- configmap.whitetail.yaml
- deployment.yaml
- service.yaml
- ingress.yaml

With PVCs

- configmap.ceres.yaml
- configmap.whitetail.yaml
- pvc.ceres.yaml
- pvc.whitetail.yaml
- deployment.pvc.yaml
- service.yaml
- ingress.yaml

No PVCs with Branding

- configmap.ceres.yaml
- configmap.whitetail.yaml
- configmap.whitetail.icon.yaml
- configmap.whitetail.logo.yaml
- deployment.branding.yaml
- service.yaml
- ingress.yaml

With PVCs and Branding

- configmap.ceres.yaml
- configmap.whitetail.yaml
- configmap.whitetail.icon.yaml
- configmap.whitetail.logo.yaml
- pvc.ceres.yaml
- pvc.whitetail.yaml
- deployment.branding.pvc.yaml
- service.yaml
- ingress.yaml

You can then send logs to ``< whitetail url >:9001`` (TCP) and ``< whitetail url >:9003`` (UDP) and can connect to the UI at ``< whitetail URL >:9001``

Contribute
----------

- Issue Tracker: github.com/jfcarter2358/whitetail/issues
- Source Code: github.com/jfcarter2358/whitetail

Support
-------

If you are having any issues please create an issue on GitHub_ or send an email to jfcarter2358@gmail.com

.. _GitHub: https://github.com/jfcarter2358/whitetail

License
-------

The project is licensed under the MIT license.