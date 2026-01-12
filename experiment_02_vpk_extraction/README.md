# Experiment 02: VPK Extraction with Overpack

## Overview

**Repository:** [gorkaerana/overpack](https://github.com/gorkaerana/overpack)  
**Type:** Python library for Veeva Vault VPK (Vault Package) files  
**Viability:** HIGH  
**Value:** HIGH - Can parse real VPK packages from SDK samples

## What Was Found

### VPK File Structure

A VPK file is a ZIP archive with this structure:

```
package.vpk
├── vaultpackage.xml          # Package manifest
├── components/               # Configuration & data components
│   ├── 00010/               # Numbered folders (execution order)
│   │   ├── ComponentType.component_name.mdl
│   │   ├── ComponentType.component_name.md5
│   │   └── ComponentType.component_name.dep  (optional)
│   ├── 00020/
│   │   ├── Label.csv        # Data component
│   │   └── Label.xml        # Data manifest
│   └── ...
└── javasdk/                  # Java SDK code (optional)
    └── src/main/java/com/veeva/vault/custom/
        └── *.java
```

### Key Classes

| Class | Purpose |
|-------|---------|
| `Vpk` | Main VPK container with manifest, components, codes |
| `ConfigurationComponent` | MDL-based configuration (Object, Picklist, etc.) |
| `DataComponent` | CSV data with manifest for import |
| `Mdl` | Wraps raw MDL with parsed `Command` via meddle |
| `Md5` | Checksum validation |
| `JavaSdkCode` | SDK source files |
| `Manifest` | Package manifest (vaultpackage.xml) |

### VPK Files Found in SDK Samples

```
experiment_06_sdk_samples/
├── vsdk-common-services-sample/
│   └── Vault-Java-SDK-Common-Services-Sample.vpk
├── vsdk-object-sample/deploy-vpk/
│   ├── vsdk-object-sample-components.vpk
│   └── vsdk-object-sample-action-code.vpk
├── vsdk-document-sample/deploy-vpk/
│   ├── Base_vsdk-document-sample-components.vpk
│   ├── Clinical_vsdk-document-sample-components.vpk
│   ├── Multichannel_vsdk-document-sample-components.vpk
│   ├── Quality_vsdk-document-sample-components.vpk
│   ├── RIM_vsdk-document-sample-components.vpk
│   └── vsdk-document-sample-code.vpk
└── vsdk-spark-integration-rules-sample/deploy-vpk/
    ├── source-vault/*.vpk
    └── target-vault/*.vpk
```

### Usage

```python
from overpack import Vpk, ConfigurationComponent

# Load a VPK file
vpk = Vpk.load("package.vpk")

# Access manifest
print(vpk.manifest.raw)

# Iterate components
for component in vpk.components:
    if isinstance(component, ConfigurationComponent):
        print(f"{component.component_type_name}.{component.component_name}")
        
        # Access parsed MDL
        if component.mdl:
            cmd = component.mdl.command
            for attr in cmd.attributes:
                print(f"  {attr.name}: {attr.value}")
            
            # Nested components (like Field inside Object)
            for sub in cmd.components:
                print(f"  {sub.component_type_name}: {sub.component_name}")

# Access Java SDK code
for code in vpk.codes:
    print(f"Java file: {code.path}")
```

### Dependency Files (.dep)

Configuration components may include `.dep` files tracking dependencies:

```csv
"In Package","Blocking Type","Source Component Label","Source Component Name",...
"false","block__sys","Access Request","access_request__c",...
"true","cascade_create__sys","Access Request","access_request__c",...
```

**Blocking Types:**
- `block__sys` - Must exist before this component
- `cascade_create__sys` - Create together
- `ignore__sys` - Optional dependency

## Integration with Meddle

Overpack uses the `meddle` library internally:

```python
from overpack import ConfigurationComponent

# When loading a ConfigurationComponent
component = ConfigurationComponent.load(path)

# The MDL is automatically parsed via meddle
if component.mdl:
    cmd = component.mdl.command  # This is a meddle.Command
    print(cmd.component_type_name)
    print(cmd.attributes)
    print(cmd.components)  # Subcomponents
```

## Conclusion

The `overpack` library provides a robust way to:

1. **Parse VPK files** - Extract configuration and data components
2. **Access MDL content** - Get parsed MDL structures via meddle
3. **Extract dependencies** - Understand component relationships
4. **Handle Java SDK code** - Access custom code files

Combined with the VPK files from the SDK samples, this allows extracting real-world Vault object definitions including fields, relationships, lifecycles, and workflows.
