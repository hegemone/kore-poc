module Kore
  module Machinery
    class MasterDeps
      include InjectMasterDeps[:log, :config]
    end
  end
end
