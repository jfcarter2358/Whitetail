AQL
==================================

.. toctree::
   :maxdepth: 2
   :caption: Contents:

========

AQL is a simple query language designed to be used when the standrard filtering (level and service) is not sufficient.  AQL statments are written in nested blocks of binary operations. This means that each operator can only have a singular left and singular right argument. An example query which gets logs of level `INFO` and level `WARN` is as follows:

``level = INFO OR level = WARN``

If you want to change the 'OR' statements to include more than just the two levels, you'll wrap the first two up in parenthesis and then OR that with a third filter.

``( level = INFO OR level = WARN ) OR level = DEBUG``

Filters
-------

The various filters that can be used in AQL statements are as follows ( ``< text like this is a placeholder >`` )

+---------------------------+-----------------------------------------------------------------------------------------------------------------------------------+
| Filter                    | Desrciption                                                                                                                       |
+===========================+===================================================================================================================================+
| level = < level >         | Get logs with level ``< level >`` (``= < level >`` can be replaced with ``IN < csv of levels >``)                                 |
+---------------------------+-----------------------------------------------------------------------------------------------------------------------------------+
| service = < service >     | Get logs from service ``< service >`` (``= < service >`` can be replaced with ``IN < csv of services >``)                         |
+---------------------------+-----------------------------------------------------------------------------------------------------------------------------------+
| year = < year >           | Get logs with a timestamp that has the year ``< year >`` (``=`` can be replaced with ``<``, ``<=``, ``>=``, ``>``, or ``!=``)     |
+---------------------------+-----------------------------------------------------------------------------------------------------------------------------------+
| month = < month >         | Get logs with a timestamp that has the month ``< month >`` (``=`` can be replaced with ``<``, ``<=``, ``>=``, ``>``, or ``!=``)   |
+---------------------------+-----------------------------------------------------------------------------------------------------------------------------------+
| day = < day >             | Get logs with a timestamp that has the day ``< day >`` (``=`` can be replaced with ``<``, ``<=``, ``>=``, ``>``, or ``!=``)       |
+---------------------------+-----------------------------------------------------------------------------------------------------------------------------------+
| hour = < hour >           | Get logs with a timestamp that has the hour ``< hour >`` (``=`` can be replaced with ``<``, ``<=``, ``>=``, ``>``, or ``!=``)     |
+---------------------------+-----------------------------------------------------------------------------------------------------------------------------------+
| minute = < minute >       | Get logs with a timestamp that has the minute ``< minute >`` (``=`` can be replaced with ``<``, ``<=``, ``>=``, ``>``, or ``!=``) |
+---------------------------+-----------------------------------------------------------------------------------------------------------------------------------+
| second = < second >       | Get logs with a timestamp that has the second ``< second >`` (``=`` can be replaced with ``<``, ``<=``, ``>=``, ``>``, or ``!=``) |
+---------------------------+-----------------------------------------------------------------------------------------------------------------------------------+
| timestamp = < timestamp > | Get logs with string timestamp in format ``YYYY-MM-DDThh:mm:ss``                                                                  |
+---------------------------+-----------------------------------------------------------------------------------------------------------------------------------+

Operators
---------

The various operators are shown below with examples

- **AND**
    - ``< left filter > AND < right filter >``
    - Returns logs which satisfy both the left and right filter
- **OR**
    - ``< left filter > OR < right filter >``
    - Returns logs which satify the left or right filter
- **NOT**
    - ``< left filter > NOT < right filter >``
    - Returns logs from the left filter that are not part of the right filter
- **XOR**
    - ``< left filter > XOR < right filter >``
    - Returns logs that are part of the left or right filter but not both
- **LIMIT**
    - Limits the results from a filter to ``N`` log messages
    - ``< left filter > LIMIT < N >``
- **ORDERBY**
    - ``< left filter > ORDERBY < field >``
    - Orders the results of the left filter in ascending order by one of the following fields
        - level
        - service
        - text
        - timestamp
        - year
        - month
        - day
        - hour
        - minute
        - second
- **ORDERDESC**
    - ``< left filter > ORDERDESC < field >``
    - Orders the results of the left filter in descending order by one of the following fields
        - level
        - service
        - text
        - timestamp
        - year
        - month
        - day
        - hour
        - minute
        - second