module Kore
  module Comm
    class EgressMessage
      attr_reader :originator, :msg
      def initialize(originator, msg)
        @originator = originator
        @msg = msg
      end
    end
  end
end
