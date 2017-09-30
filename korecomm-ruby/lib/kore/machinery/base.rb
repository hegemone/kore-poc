module Kore
  module Machinery
    class Base
      def initialize
        # Compose base of deps to avoid polluting the ctor of children
        @deps = Kore::Machinery::MasterDeps.new
      end
      # Convenience accessors for children
      def log
        @deps.log
      end
      def config
        @deps.config
      end
    end
  end
end
